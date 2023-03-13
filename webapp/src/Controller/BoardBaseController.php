<?php

namespace App\Controller;

use Doctrine\Persistence\ObjectManager;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Mercure\HubInterface;
use Symfony\Component\Mercure\Update;
use Symfony\Component\Routing\Generator\UrlGeneratorInterface;
use Symfony\Component\Serializer\SerializerInterface;

use App\Entity\Board;
use App\Entity\User;
use App\Service\Direction;
use App\ViewModel\BoardViewModel;
use App\ViewModel\PlayerViewModel;
use App\Service\DomainServiceInterface;
use App\Service\Rotation;

abstract class BoardBaseController extends AbstractController
{
    const SESSION_PLAYER_KEY = 'player';

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

        foreach ($board->getUsers() as $gamePlayer) {
            if ($gamePlayer->getId() == $user->getId()) {
                return false;
            }
        }
        return true;
    }

    protected function createBoardViewModel(?User $user, Board $board): BoardViewModel
    {
        $state = $board->getState();
        $boardPlayers = $board->getUsers();
        $players =
            array_map(
                function ($player, $index) use ($state, $user, $boardPlayers) {
                    $boardUser = $boardPlayers[$index];
                    if (!$boardUser) {
                        return null;
                    }

                    $isCurrentPlayer =
                        count($state['remainingPlayers']) > 0 &&
                        $state['remainingPlayers'][$state['currentPlayerIndex']] == $index;
                    $isUser = $user && $boardUser->getId() == $user->getId();

                    return new PlayerViewModel(
                        $boardUser->getUsername(),
                        $player['color'],
                        $player['position']['line'],
                        $player['position']['row'],
                        $player['targets'],
                        $player['score'],
                        $isCurrentPlayer,
                        $isUser
                    );
                },
                $state['players'],
                array_keys($state['players'])
            );

        $canPlay = $user != null && $state['gameState'] != 2 && current(array_filter($players, function ($player) {
            return $player && $player->getIsCurrentPlayer() && $player->getIsUser();
        })) !== false;

        return new BoardViewModel(
            $board->getId(),
            $board->getRemainingSeats(),
            $this->canUserJoin($user, $board),
            $state,
            $players,
            $canPlay,
        );
    }

    protected function canUserPlay(?User $user, Board $board): bool
    {
        return $this->createBoardViewModel($user, $board)->getCanPlay();
    }

    protected function publishUpdate(Board $board)
    {
        $update = new Update(
            $this->generateUrl('board_view', ['id' => $board->getId()], UrlGeneratorInterface::ABSOLUTE_URL),
            $this->serializer->serialize([], 'json')
        );

        $this->hub->publish($update);
    }

    protected function rotateRemaining(Board $board, Rotation $rotation)
    {
        $updatedBoard = $this->domainService->rotateRemainingTile($board->getState(), $rotation);
        $board->setState($updatedBoard);

        $this->entityManager->flush();
        $this->publishUpdate($board);
    }

    protected function insertTile(Board $board, Direction $direction, int $index)
    {
        $updatedBoard = $this->domainService->insertTile(
            $board->getState(),
            $direction,
            $index,
        );
        $board->setState($updatedBoard);

        $this->entityManager->flush();
        $this->publishUpdate($board);
    }

    protected function movePlayer(Board $board, int $line, int $row)
    {
        $updatedBoard = $this->domainService->movePlayer(
            $board->getState(),
            $line,
            $row,
        );
        $board->setState($updatedBoard);

        $this->entityManager->flush();
        $this->publishUpdate($board);
    }

    protected function newBoard(User $user, int $playerCount): Board
    {
        $boardState = $this->domainService->newBoard(intval($playerCount));

        $board = new Board();
        $board->setState($boardState);
        $board->addUser($user);
        $board->setRemainingSeats($playerCount - 1);

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

            $board->setRemainingSeats($remainingSeats - 1);
            $board->addUser($user);

            $entityManager->flush();
            $conn->commit();

            $this->publishUpdate($board);
            return true;
        } catch (\Exception $e) {
            $conn->rollBack();
            throw $e;
        }
    }
}
