<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230316080326 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        $this->addSql("
INSERT INTO \"user\"(\"id\", \"username\", \"email\", \"password\", \"roles\", \"created_at\", \"updated_at\")
VALUES (nextval('user_id_seq'), 'admin', 'admin@marmelab.com', '\$2y\$13\$XKVM2eNNBDWdvh9lUzApxuR1G.cHRO80bfJSBXjfXSEPFxzSp.E7y', '[\"ROLE_ADMIN\"]', NOW(), NOW())");
    }

    public function down(Schema $schema): void
    {
        $this->addSql("DELETE FROM \"user\" WHERE \"username\" = 'admin'");
    }
}
