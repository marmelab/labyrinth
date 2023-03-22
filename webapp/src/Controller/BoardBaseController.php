<?php

namespace App\Controller;

use Doctrine\Persistence\ObjectManager;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\Mercure\HubInterface;
use Symfony\Component\Mercure\Update;
use Symfony\Component\Routing\Generator\UrlGeneratorInterface;
use Symfony\Component\Serializer\SerializerInterface;

use App\Entity\Board;
use App\Entity\Player;
use App\Entity\User;
use App\Service\Direction;
use App\ViewModel\AccessibleTilesViewModel;
use App\ViewModel\BoardViewModel;
use App\ViewModel\PlayerViewModel;
use App\Service\DomainServiceInterface;
use App\Service\Rotation;

abstract class BoardBaseController extends AbstractController
{
    const GAME_STATE_PLACE_TILE = 0;
    const GAME_STATE_MOVE_PAWN = 1;
    const GAME_STATE_END = 2;

    protected function __construct(
        protected DomainServiceInterface $domainService,
        protected ObjectManager $entityManager,
        protected HubInterface $hub,
        protected SerializerInterface $serializer,
    ) {
    }

    protected function canUserJoin(?User $user, Board $board): bool
    {
        if ($user == null) {
            return false;
        }

        foreach ($board->getPlayers() as $player) {
            if ($player->isIsBot()) {
                return false;
            }

            if ($player->getAttendee()->getId() == $user->getId()) {
                return false;
            }
        }
        return true;
    }

    protected function createBoardViewModel(?User $user, Board $board): BoardViewModel
    {
        $state = $board->getState();
        $players =
            array_map(
                /** @var Player $player */
                function ($player) use ($user) {
                    $attendee = $player->getAttendee();


                    $isUser = $user && $attendee && $attendee->getId() == $user->getId();

                    return new PlayerViewModel(
                        $attendee ? $attendee->getUsername() : "Bot #" . $player->getId(),
                        $player->isIsBot(),
                        $player->getColor(),
                        $player->getLine(),
                        $player->getRow(),
                        $player->getTargets(),
                        $player->getScore(),
                        $player->isCurrentPlayer(),
                        $isUser
                    );
                },
                $board->getPlayers()->toArray()
            );

        $canPlay = $user && $state['gameState'] != static::GAME_STATE_END && current(array_filter($players, function ($player) {
            return $player && $player->getIsCurrentPlayer() && $player->getIsUser();
        })) !== false;

        $isGameCreator = $user && $state['gameState'] != static::GAME_STATE_END && current(array_filter($players, function ($player) {
            return $player && $player->getIsUser();
        })) !== false;

        /** @var ?AccessibleTilesViewModel */
        $accessibleTiles = null;
        if ($canPlay && $state['gameState'] == static::GAME_STATE_MOVE_PAWN) {
            $accessibleTilesData = $this->domainService->getAccessibleTiles($state);
            $accessibleTiles = new AccessibleTilesViewModel(
                $accessibleTilesData['isShortestPath'],
                $accessibleTilesData['coordinates']
            );
        }

        return new BoardViewModel(
            $board->getId(),
            $board->getRemainingSeats(),
            $this->canUserJoin($user, $board),
            $state,
            $players,
            $canPlay,
            $isGameCreator,
            $accessibleTiles,
        );
    }

    protected function canUserPlay(?User $user, Board $board): bool
    {
        $boardViewModel = $this->createBoardViewModel($user, $board);
        if ($boardViewModel->getCanPlay()) {
            return true;
        }

        if ($boardViewModel->getCurrentPlayer()->getIsBot() && $boardViewModel->getIsGameCreator()) {
            return true;
        }

        return false;
    }


    protected function publishUpdate(Board $board, array $actions)
    {
        $update = new Update(
            $this->generateUrl('board_view', ['id' => $board->getId()]),
            $this->serializer->serialize($actions, 'json')
        );

        $this->hub->publish($update);
    }


    protected function bindPlayerFromState(Player $player, array $state, int $playerIndex): Player
    {
        $remainingPlayers = $state['remainingPlayers'];
        $isCurrentPlayer =
            count($remainingPlayers) > 0 &&
            $remainingPlayers[$state['currentPlayerIndex']] == $playerIndex;


        usort($state['players'], function ($a, $b) {
            return $a['color'] <=> $b['color'];
        });

        $playerState = $state['players'][$playerIndex];
        return $player
            ->setColor($playerState['color'])
            ->setTargets($playerState['targets'])
            ->setScore($playerState['score'])
            ->setLine($playerState['position']['line'])
            ->setRow($playerState['position']['row'])
            ->setCurrentPlayer($isCurrentPlayer);
    }

    protected function updateBoard(Board $board, array $newState)
    {
        $board->setState($newState);
        $board->setGameState($newState['gameState'])->setUpdatedAt();

        $players = $board->getPlayers()->toArray();

        $maxWinOrder = array_reduce($players, function ($current, $player) {
            /** @var Player $player */

            $winOrder = $player->getWinOrder();
            if ($winOrder != null && $winOrder > $current) {
                return $winOrder;
            }
            return $current;
        }, 0);

        usort($players, function ($a, $b) {
            return $a->getColor() <=> $b->getColor();
        });
        foreach ($players as $index => $player) {
            /** @var Player $player */

            $this->bindPlayerFromState($player, $board->getState(), $index);
            if ($player->getWinOrder() == null && count($player->getTargets()) == 0) {
                $player->setWinOrder(++$maxWinOrder);
            }
        }
    }


    protected function rotateRemaining(Board $board, Rotation $rotation)
    {
        $newState = $this->domainService->rotateRemainingTile($board->getState(), $rotation);
        $this->updateBoard($board, $newState['board']);

        $this->entityManager->flush();
        $this->publishUpdate($board, $newState['actions']);
    }

    protected function insertTile(Board $board, Direction $direction, int $index)
    {
        $newState = $this->domainService->insertTile(
            $board->getState(),
            $direction,
            $index,
        );
        $this->updateBoard($board, $newState['board']);


        $this->entityManager->flush();
        $this->publishUpdate($board, $newState['actions']);
    }

    protected function movePlayer(Board $board, int $line, int $row)
    {
        $newState = $this->domainService->movePlayer(
            $board->getState(),
            $line,
            $row,
        );
        $this->updateBoard($board, $newState['board']);

        $this->entityManager->flush();
        $this->publishUpdate($board, $newState['actions']);
    }

    protected function newBoard(User $user, int $playerCount, bool $isBotGame = false): Board
    {
        $boardState = $this->domainService->newBoard(intval($playerCount));

        $player = $this->bindPlayerFromState(new Player(), $boardState, 0);
        $player
            ->setIsBot(false)
            ->setAttendee($user);
        $this->entityManager->persist($player);

        $board = (new Board())
            ->setState($boardState)
            ->setGameState(0)
            ->addPlayer($player)
            ->setRemainingSeats($isBotGame ? 0 : $playerCount - 1)
            ->setCreatedAt()
            ->setUpdatedAt();

        if ($isBotGame) {
            for ($i = 1; $i < $playerCount; $i++) {
                $player = $this->bindPlayerFromState(new Player(), $boardState, $i);
                $player
                    ->setIsBot(true)
                    ->setAttendee(null);
                $this->entityManager->persist($player);
                $board
                    ->addPlayer($player);
            }
        }

        $this->entityManager->persist($board);
        $this->entityManager->flush();

        return $board;
    }

    protected function joinBoard(User $user, Board $board): bool
    {
        if (!$this->canUserJoin($user, $board)) {
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        }

        /** @var Doctrine\Persistence\EntityManager $entityManager */
        $entityManager = $this->entityManager;

        $conn = $entityManager->getConnection();
        $conn->beginTransaction();
        $conn->setAutoCommit(false);
        try {
            if (!$this->canUserJoin($user, $board)) {
                $conn->rollBack();
                return false;
            }

            // This will ensure that no concurrent updates happen on the board.
            $boardRepository = $entityManager->getRepository(Board::class);

            /** @var Board $board */
            $board = $boardRepository->find($board->getId());

            $remainingSeats = $board->getRemainingSeats();
            if ($remainingSeats == 0) {
                $conn->rollBack();
                return false;
            }

            $playerCount = $board->getPlayers()->count();
            $player = $this->bindPlayerFromState(new Player(), $board->getState(), $playerCount);
            $player
                ->setIsBot(false)
                ->setAttendee($user);
            $entityManager->persist($player);

            $board->setRemainingSeats($remainingSeats - 1);
            $board->addPlayer($player);

            $entityManager->flush();
            $conn->commit();

            $this->publishUpdate($board, [[
                'kind' => 'NEW_PLAYER',
                'payload' => $player,
            ]]);
            return true;
        } catch (\Exception $e) {
            $conn->rollBack();
            throw $e;
        }
    }
}
