<?php

namespace App\Repository;

use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

use App\Entity\Board;
use App\Entity\User;


/**
 * @extends ServiceEntityRepository<Board>
 *
 * @method Board|null find($id, $lockMode = null, $lockVersion = null)
 * @method Board|null findOneBy(array $criteria, array $orderBy = null)
 * @method Board[]    findAll()
 * @method Board[]    findBy(array $criteria, array $orderBy = null, $limit = null, $offset = null)
 */
class BoardRepository extends ServiceEntityRepository
{
    const PAGE_SIZE = 25;

    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Board::class);
    }

    public function save(Board $entity, bool $flush = false): void
    {
        $this->getEntityManager()->persist($entity);

        if ($flush) {
            $this->getEntityManager()->flush();
        }
    }

    public function remove(Board $entity, bool $flush = false): void
    {
        $this->getEntityManager()->remove($entity);

        if ($flush) {
            $this->getEntityManager()->flush();
        }
    }

    public function find($id, $lockMode = null, $lockVersion = null): ?Board
    {
        $qb = $this->createQueryBuilder('b')
            ->select(['b', 'p', 'u'])
            ->leftJoin('b.players', 'p')
            ->leftJoin('p.attendee', 'u')
            ->where('b.id = :boardId')
            ->orderBy('p.color')
            ->setParameter('boardId', $id);

        return $qb->getQuery()->getSingleResult();
    }

    public function findByAnonymous(int $page = 1, int $pageSize = 25): array
    {
        $qb = $this->createQueryBuilder('b')
            ->select(['b.id', 'b.remainingSeats'])
            ->setFirstResult(($page - 1) * $pageSize)
            ->setMaxResults($pageSize);

        return $qb->getQuery()->execute();
    }

    public function findByUser(User $user, int $page = 1, int $pageSize = 25): array
    {
        $qb = $this->createQueryBuilder('b')
            ->select(['b.id', 'b.remainingSeats'])
            ->leftJoin('b.players', 'p')
            ->leftJoin('p.attendee', 'a')
            ->where('a.id = :userId')
            ->setParameter('userId', $user->getId())
            ->orderBy('b.id', 'DESC')
            ->setFirstResult(($page - 1) * $pageSize)
            ->setMaxResults($pageSize);

        return $qb->getQuery()->execute();
    }
}
