<?php

namespace App\ViewModel;

use Symfony\Component\Serializer\Annotation\SerializedName;

class BoardViewModel
{
    public function __construct(
        private int $id,
        private int $remainingSeats,
        private bool $canJoin,
        private array $state,
        private array $players,
        private bool $canPlay,
        private bool $isGameCreator,
        private ?LastInsertionViewModel $lastInsertion,
        private ?AccessibleTilesViewModel $accessibleTiles = null
    ) {
    }

    public function getId(): int
    {
        return $this->id;
    }

    public function getRemainingSeats(): int
    {
        return $this->remainingSeats;
    }

    public function getCanJoin(): bool
    {
        return $this->canJoin;
    }


    public function getState(): array
    {
        return $this->state;
    }

    public function getPlayers(): array
    {
        return $this->players;
    }

    #[SerializedName("currentPlayer")]
    public function getCurrentPlayer(): ?PlayerViewModel
    {
        foreach ($this->players as $player) {
            if ($player && $player->getIsCurrentPlayer()) {
                return $player;
            }
        }
        return null;
    }

    #[SerializedName("user")]
    public function getUser(): ?PlayerViewModel
    {
        foreach ($this->players as $player) {
            if ($player && $player->getIsUser()) {
                return $player;
            }
        }
        return null;
    }

    public function getCanPlay(): bool
    {
        return $this->canPlay;
    }

    #[SerializedName("gameState")]
    public function getGameState(): int
    {
        return $this->state['gameState'];
    }

    public function getLastInsertion(): ?LastInsertionViewModel
    {
        return $this->lastInsertion;
    }

    public function getAccessibleTiles(): ?AccessibleTilesViewModel
    {
        return $this->accessibleTiles;
    }

    public function getIsGameCreator(): bool
    {
        return $this->isGameCreator;
    }
}
