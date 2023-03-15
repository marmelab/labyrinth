<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230315085234 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        $this->addSql('GRANT SELECT, INSERT, UPDATE, DELETE ON "board_user" TO role_admin');
    }

    public function down(Schema $schema): void
    {
        $this->addSql('REVOKE ALL ON "board_user" FROM role_admin');
    }
}
