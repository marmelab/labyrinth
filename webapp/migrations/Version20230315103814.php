<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230315103814 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        $this->addSql('TRUNCATE TABLE board RESTART IDENTITY CASCADE');

        $this->addSql('CREATE SEQUENCE player_id_seq INCREMENT BY 1 MINVALUE 1 START 1');
        $this->addSql('CREATE TABLE player (id INT NOT NULL, board_id INT NOT NULL, attendee_id INT NOT NULL, color INT NOT NULL, targets JSON NOT NULL, score INT NOT NULL, line INT NOT NULL, row INT NOT NULL, win_order INT DEFAULT NULL, PRIMARY KEY(id))');
        $this->addSql('CREATE INDEX IDX_98197A65E7EC5785 ON player (board_id)');
        $this->addSql('CREATE INDEX IDX_98197A65BCFD782A ON player (attendee_id)');
        $this->addSql('ALTER TABLE player ADD CONSTRAINT FK_98197A65E7EC5785 FOREIGN KEY (board_id) REFERENCES board (id) NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE player ADD CONSTRAINT FK_98197A65BCFD782A FOREIGN KEY (attendee_id) REFERENCES "user" (id) NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE board_user DROP CONSTRAINT fk_57058f6ae7ec5785');
        $this->addSql('ALTER TABLE board_user DROP CONSTRAINT fk_57058f6aa76ed395');
        $this->addSql('DROP TABLE board_user');
        $this->addSql('ALTER TABLE board ALTER created_at DROP DEFAULT');
        $this->addSql('ALTER TABLE board ALTER updated_at DROP DEFAULT');
        $this->addSql('ALTER TABLE "user" ALTER created_at DROP DEFAULT');
        $this->addSql('ALTER TABLE "user" ALTER updated_at DROP DEFAULT');

        $this->addSql('GRANT SELECT, INSERT, UPDATE, DELETE ON "player" TO role_admin');
    }

    public function down(Schema $schema): void
    {
        $this->addSql('REVOKE ALL ON "player" FROM role_admin');

        $this->addSql('DROP SEQUENCE player_id_seq CASCADE');
        $this->addSql('CREATE TABLE board_user (board_id INT NOT NULL, user_id INT NOT NULL, PRIMARY KEY(board_id, user_id))');
        $this->addSql('CREATE INDEX idx_57058f6aa76ed395 ON board_user (user_id)');
        $this->addSql('CREATE INDEX idx_57058f6ae7ec5785 ON board_user (board_id)');
        $this->addSql('ALTER TABLE board_user ADD CONSTRAINT fk_57058f6ae7ec5785 FOREIGN KEY (board_id) REFERENCES board (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE board_user ADD CONSTRAINT fk_57058f6aa76ed395 FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE');
        $this->addSql('ALTER TABLE player DROP CONSTRAINT FK_98197A65E7EC5785');
        $this->addSql('ALTER TABLE player DROP CONSTRAINT FK_98197A65BCFD782A');
        $this->addSql('DROP TABLE player');
        $this->addSql('ALTER TABLE board ALTER created_at SET DEFAULT \'now()\'');
        $this->addSql('ALTER TABLE board ALTER updated_at SET DEFAULT \'now()\'');
        $this->addSql('ALTER TABLE "user" ALTER created_at SET DEFAULT \'now()\'');
        $this->addSql('ALTER TABLE "user" ALTER updated_at SET DEFAULT \'now()\'');
    }
}
