<?php

namespace App\Tests\Unit\Controller;

use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

use App\Service\DomainServiceInterface;

class HomeControllerTest extends WebTestCase
{
    public function provideBoardData(): array
    {
        $board = json_decode(file_get_contents("./tests/Unit/Data/board.json"), true);
        return [
            [$board],
        ];
    }

    /**
     * @dataProvider provideBoardData
     */
    function testIndex(array $board)
    {
        $client = static::createClient();

        $crawler = $client->request('GET', '/');

        $this->assertResponseIsSuccessful();
        $this->assertCount(1, $crawler->filter('button'));
    }
}
