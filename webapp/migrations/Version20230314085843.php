<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230314085843 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE TABLE board_user (board_id INT NOT NULL, user_id INT NOT NULL, PRIMARY KEY(board_id, user_id))');
        $this->addSql('CREATE INDEX IDX_57058F6AE7EC5785 ON board_user (board_id)');
        $this->addSql('CREATE INDEX IDX_57058F6AA76ED395 ON board_user (user_id)');
        $this->addSql('ALTER TABLE board_user ADD CONSTRAINT FK_57058F6AE7EC5785 FOREIGN KEY (board_id) REFERENCES board (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE board_user ADD CONSTRAINT FK_57058F6AA76ED395 FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE board_player DROP CONSTRAINT fk_34073ec2e7ec5785');
        $this->addSql('ALTER TABLE board_player DROP CONSTRAINT fk_34073ec299e6f5df');
        $this->addSql('DROP TABLE board_player');
        $this->addSql('TRUNCATE TABLE board RESTART IDENTITY CASCADE');
    }

    public function down(Schema $schema): void
    {
        // this down() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE SCHEMA public');
        $this->addSql('CREATE TABLE board_player (board_id INT NOT NULL, player_id INT NOT NULL, PRIMARY KEY(board_id, player_id))');
        $this->addSql('CREATE INDEX idx_34073ec299e6f5df ON board_player (player_id)');
        $this->addSql('CREATE INDEX idx_34073ec2e7ec5785 ON board_player (board_id)');
        $this->addSql('ALTER TABLE board_player ADD CONSTRAINT fk_34073ec2e7ec5785 FOREIGN KEY (board_id) REFERENCES board (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE board_player ADD CONSTRAINT fk_34073ec299e6f5df FOREIGN KEY (player_id) REFERENCES player (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE board_user DROP CONSTRAINT FK_57058F6AE7EC5785');
        $this->addSql('ALTER TABLE board_user DROP CONSTRAINT FK_57058F6AA76ED395');
        $this->addSql('DROP TABLE board_user');
    }
}
