<?php

namespace App\Entity;

use DateTime;

use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;

use App\Repository\BoardRepository;

#[ORM\Entity(repositoryClass: BoardRepository::class)]
class Board
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column(type: Types::JSON)]
    private ?array $state = null;

    #[ORM\Column(options: ["default" => 0])]
    private ?int $remainingSeats = null;

    #[ORM\Column(type: 'datetime')]
    private DateTime $createdAt;

    #[ORM\Column(type: 'datetime')]
    private DateTime $updatedAt;

    #[ORM\OneToMany(mappedBy: 'board', targetEntity: Player::class, cascade: ["all"])]
    private Collection $players;

    #[ORM\Column]
    private ?int $gameState = null;

    public function __construct()
    {
        $this->players = new ArrayCollection();
    }

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getState(): ?array
    {
        return $this->state;
    }

    public function setState(array $state): self
    {
        $this->state = $state;

        return $this;
    }

    public function getRemainingSeats(): ?int
    {
        return $this->remainingSeats;
    }

    public function setRemainingSeats(int $remainingSeats): self
    {
        $this->remainingSeats = $remainingSeats;

        return $this;
    }

    public function getCreatedAt(): DateTime
    {
        return $this->createdAt;
    }

    public function setCreatedAt(): self
    {
        $this->createdAt = new DateTime("now");
        return $this;
    }

    public function getUpdatedAt(): DateTime
    {
        return $this->updatedAt;
    }

    public function setUpdatedAt(): self
    {
        $this->updatedAt = new DateTime("now");

        return $this;
    }

    /**
     * @return Collection<int, Player>
     */
    public function getPlayers(): Collection
    {
        return $this->players;
    }

    public function addPlayer(Player $player): self
    {
        if (!$this->players->contains($player)) {
            $this->players->add($player);
            $player->setBoard($this);
        }

        return $this;
    }

    public function removePlayer(Player $player): self
    {
        if ($this->players->removeElement($player)) {
            // set the owning side to null (unless already changed)
            if ($player->getBoard() === $this) {
                $player->setBoard(null);
            }
        }

        return $this;
    }

    public function getGameState(): ?int
    {
        return $this->gameState;
    }

    public function setGameState(int $gameState): self
    {
        $this->gameState = $gameState;

        return $this;
    }
}
