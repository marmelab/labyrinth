<?php

namespace App\ViewModel;

use Symfony\Component\Serializer\Annotation\SerializedName;

class PlayerViewModel
{

    public function __construct(
        private string $name,
        private bool $isBot,
        private int $color,
        private int $line,
        private int $row,
        private array $targets,
        private int $score,
        private ?int $winOrder,
        private bool $isCurrentPlayer,
        private bool $isUser,
    ) {
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function getIsBot(): bool
    {
        return $this->isBot;
    }

    public function getColor(): int
    {
        return $this->color;
    }

    public function getLine(): int
    {
        return $this->line;
    }

    public function getRow(): int
    {
        return $this->row;
    }

    public function getTargets(): array
    {
        return $this->targets;
    }

    public function getScore(): int
    {
        return $this->score;
    }

    public function getWinOrder(): ?int
    {
        return $this->winOrder;
    }

    #[SerializedName("totalTargets")]
    public function getTotalTargets(): int
    {
        return $this->score + count($this->targets);
    }


    #[SerializedName("currentTarget")]
    public function getCurrentTarget(): string
    {
        if (count($this->targets) > 0) {
            return $this->targets[0];
        }
        return '';
    }

    public function getIsCurrentPlayer(): bool
    {
        return $this->isCurrentPlayer;
    }

    public function getIsUser(): bool
    {
        return $this->isUser;
    }
}
