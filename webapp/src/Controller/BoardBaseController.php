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
use App\Entity\Player;
use App\Service\Direction;
use App\ViewModel\BoardViewModel;
use App\ViewModel\PlayerViewModel;
use App\Service\DomainServiceInterface;
use App\Service\Rotation;

abstract class BoardBaseController extends AbstractController
{
    const SESSION_PLAYER_KEY = AuthBaseController::SESSION_PLAYER_KEY;

    protected function __construct(
        protected DomainServiceInterface $domainService,
        protected ObjectManager $entityManager,
        protected HubInterface $hub,
        protected SerializerInterface $serializer,
    ) {
    }

    protected function getCurrentUser(Request $request): ?Player
    {
        $player = $request->getSession()->get(static::SESSION_PLAYER_KEY);
        if ($player == NULL) {
            return NULL;
        }

        $playerRepository = $this->entityManager->getRepository(Player::class);
        return $playerRepository->find($player->getId());
    }

    protected function canUserJoin(?Player $user, Board $board): bool
    {
        if ($user == null) {
            return false;
        }

        foreach ($board->getPlayers() as $gamePlayer) {
            if ($gamePlayer->getId() == $user->getId()) {
                return false;
            }
        }
        return true;
    }

    protected function createBoardViewModel(?Player $user, Board $board): BoardViewModel
    {
        $state = $board->getState();
        $boardPlayers = $board->getPlayers();
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
                        $boardUser->getName(),
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
            $state,
            $players,
            $canPlay,
        );
    }

    protected function canUserPlay(?Player $user, Board $board): bool
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
}
