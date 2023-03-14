<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230314143326 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        $this->addSql('REVOKE ALL ON "board" FROM anonymous');
        $this->addSql('REVOKE ALL ON "user" FROM anonymous');
    }

    public function down(Schema $schema): void
    {
        $this->addSql('GRANT SELECT, INSERT, UPDATE, DELETE ON "board" TO anonymous');
        $this->addSql('GRANT SELECT, INSERT, UPDATE, DELETE ON "user" TO anonymous');
    }
}
