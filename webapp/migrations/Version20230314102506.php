<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230314102506 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        $this->addSql('CREATE ROLE ROLE_ADMIN NOLOGIN');
        $this->addSql('GRANT SELECT, INSERT, UPDATE, DELETE ON "board" TO ROLE_ADMIN');
        $this->addSql('GRANT SELECT, INSERT, UPDATE, DELETE ON "user" TO ROLE_ADMIN');

        $this->addSql('GRANT SELECT, INSERT, UPDATE, DELETE ON "board" TO anonymous');
        $this->addSql('GRANT SELECT, INSERT, UPDATE, DELETE ON "user" TO anonymous');
    }

    public function down(Schema $schema): void
    {
        $this->addSql('REVOKE ALL ON "board" FROM anonymous');
        $this->addSql('REVOKE ALL ON "user" FROM anonymous');

        $this->addSql('REVOKE ALL ON "board" FROM ROLE_ADMIN');
        $this->addSql('REVOKE ALL ON "user" FROM ROLE_ADMIN');
        $this->addSql('DROP ROLE ROLE_ADMIN');
    }
}
