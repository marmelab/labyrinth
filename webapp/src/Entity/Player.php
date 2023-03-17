<?php

namespace App\Entity;

use Doctrine\ORM\Mapping as ORM;
use Symfony\Component\Serializer\Annotation\Ignore;

use App\Repository\PlayerRepository;

#[ORM\Entity(repositoryClass: PlayerRepository::class)]
class Player
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column]
    private ?int $color = null;

    #[ORM\Column]
    private array $targets = [];

    #[ORM\Column]
    private ?int $score = null;

    #[ORM\Column]
    private ?int $line = null;

    #[ORM\Column]
    private ?int $row = null;

    #[ORM\Column(nullable: true)]
    private ?int $winOrder = null;

    #[ORM\ManyToOne(inversedBy: 'players')]
    #[ORM\JoinColumn(nullable: false)]
    #[Ignore]
    private ?Board $board = null;

    #[ORM\ManyToOne(inversedBy: 'games')]
    #[ORM\JoinColumn(nullable: false)]
    private ?User $attendee = null;

    #[ORM\Column]
    private ?bool $currentPlayer = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getColor(): ?int
    {
        return $this->color;
    }

    public function setColor(int $color): self
    {
        $this->color = $color;

        return $this;
    }

    public function getTargets(): array
    {
        return $this->targets;
    }

    public function setTargets(array $targets): self
    {
        $this->targets = $targets;

        return $this;
    }

    public function getScore(): ?int
    {
        return $this->score;
    }

    public function setScore(int $score): self
    {
        $this->score = $score;

        return $this;
    }

    public function getLine(): ?int
    {
        return $this->line;
    }

    public function setLine(int $line): self
    {
        $this->line = $line;

        return $this;
    }

    public function getRow(): ?int
    {
        return $this->row;
    }

    public function setRow(int $row): self
    {
        $this->row = $row;

        return $this;
    }

    public function getWinOrder(): ?int
    {
        return $this->winOrder;
    }

    public function setWinOrder(?int $winOrder): self
    {
        $this->winOrder = $winOrder;

        return $this;
    }

    public function getBoard(): ?Board
    {
        return $this->board;
    }

    public function setBoard(?Board $board): self
    {
        $this->board = $board;

        return $this;
    }

    public function getAttendee(): ?User
    {
        return $this->attendee;
    }

    public function setAttendee(?User $attendee): self
    {
        $this->attendee = $attendee;

        return $this;
    }

    public function isCurrentPlayer(): ?bool
    {
        return $this->currentPlayer;
    }

    public function setCurrentPlayer(bool $currentPlayer): self
    {
        $this->currentPlayer = $currentPlayer;

        return $this;
    }
}
