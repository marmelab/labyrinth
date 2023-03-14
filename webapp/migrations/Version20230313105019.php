<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230313105019 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        $this->addSql('ALTER TABLE board ALTER remaining_seats SET DEFAULT 0');
        $this->addSql('CREATE ROLE anonymous NOLOGIN;');
        $this->addSql('CREATE ROLE authenticator NOLOGIN NOINHERIT;');
    }

    public function down(Schema $schema): void
    {
        $this->addSql('ALTER TABLE board ALTER remaining_seats DROP DEFAULT');
        $this->addSql('DROP ROLE anonymous;');
        $this->addSql('DROP ROLE authenticator;');
    }
}
