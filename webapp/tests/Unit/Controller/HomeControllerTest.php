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

        $container = static::getContainer();

        $domainServiceMock = $this->createMock(DomainServiceInterface::class);
        $domainServiceMock->expects(self::once())
            ->method('newBoard')
            ->willReturn($board);

        $container->set(DomainServiceInterface::class, $domainServiceMock);

        $crawler = $client->request('GET', '/');

        $this->assertResponseIsSuccessful();
        $this->assertCount(1, $crawler->filter('.board'));
        $this->assertCount(49, $crawler->filter('.board .tile'));
    }
}