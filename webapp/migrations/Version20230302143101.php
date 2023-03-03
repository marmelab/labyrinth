<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230302143101 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE TABLE board_player (board_id INT NOT NULL, player_id INT NOT NULL, PRIMARY KEY(board_id, player_id))');
        $this->addSql('CREATE INDEX IDX_34073EC2E7EC5785 ON board_player (board_id)');
        $this->addSql('CREATE INDEX IDX_34073EC299E6F5DF ON board_player (player_id)');
        $this->addSql('ALTER TABLE board_player ADD CONSTRAINT FK_34073EC2E7EC5785 FOREIGN KEY (board_id) REFERENCES board (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE board_player ADD CONSTRAINT FK_34073EC299E6F5DF FOREIGN KEY (player_id) REFERENCES player (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE');
    }

    public function down(Schema $schema): void
    {
        // this down() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE SCHEMA public');
        $this->addSql('ALTER TABLE board_player DROP CONSTRAINT FK_34073EC2E7EC5785');
        $this->addSql('ALTER TABLE board_player DROP CONSTRAINT FK_34073EC299E6F5DF');
        $this->addSql('DROP TABLE board_player');
    }
}
