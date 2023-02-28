<?php

namespace App\Tests\Unit\Service;

use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;
use Symfony\Component\HttpClient\MockHttpClient;
use Symfony\Component\HttpClient\Response\MockResponse;
use Symfony\Contracts\HttpClient\Exception\ServerExceptionInterface;

use App\Service\DomainService;

class DomainServiceTest extends KernelTestCase
{
    public function provideBoardData(): array
    {
        $boardJson = file_get_contents("./tests/Unit/Data/board.json");
        $board = json_decode($boardJson, true);
        return [
            [$boardJson, $board],
        ];
    }

    /**
     * @dataProvider provideBoardData
     */
    public function testNewBoard__Ok(string $boardJson, array $board): void
    {
        $mockResponse = new MockResponse($boardJson);
        $mockClient = new MockHttpClient($mockResponse);

        $domainServiceClient = new DomainService(
            $mockClient,
            "http://domain-api"
        );

        $this->assertEquals($board, $domainServiceClient->newBoard());
    }

    public function testNewBoard__InternalServererror(): void
    {
        $mockResponse = new MockResponse("", [
            'http_code' => 500,
        ]);
        $mockClient = new MockHttpClient($mockResponse);

        $domainServiceClient = new DomainService(
            $mockClient,
            "http://domain-api"
        );

        $this->expectException(ServerExceptionInterface::class);
        $domainServiceClient->newBoard();
    }
}