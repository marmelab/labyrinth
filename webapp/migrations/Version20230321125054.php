<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230321125054 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        $this->addSql('ALTER TABLE player ADD is_bot BOOLEAN NOT NULL DEFAULT false;');
        $this->addSql('ALTER TABLE player ALTER attendee_id DROP NOT NULL;');
    }

    public function down(Schema $schema): void
    {
        $this->addSql('ALTER TABLE player DROP is_bot;');
        $this->addSql('ALTER TABLE player ALTER attendee_id SET NOT NULL');
    }
}
