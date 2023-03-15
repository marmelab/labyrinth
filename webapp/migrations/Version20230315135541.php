<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230315135541 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql("CREATE VIEW ongoing_game AS
SELECT p.id, p.board_id, p.attendee_id, b.game_state 
FROM player p 
INNER JOIN board b ON p.board_id = b.id
WHERE b.game_state < 2;");
        $this->addSql('GRANT SELECT ON "ongoing_game" TO role_admin');
    }

    public function down(Schema $schema): void
    {
        $this->addSql("DROP VIEW ongoing_game;");
    }
}
