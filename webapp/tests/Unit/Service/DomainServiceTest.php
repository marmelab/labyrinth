<?php

namespace App\Tests\Unit\Service;

use App\Service\Rotation;
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

    public function provideRotateRemainingTileData(): array
    {
        $boardJson = file_get_contents("./tests/Unit/Data/board.json");
        $board = json_decode($boardJson, true);
        return [
            [
                json_encode(array_replace_recursive($board, ['remainingTile' => ['rotation' => 90]])),
                $board,
                Rotation::CLOCKWISE,
                90
            ],
            [
                json_encode(array_replace_recursive($board, ['remainingTile' => ['rotation' => 270]])),
                $board,
                Rotation::ANTICLOCKWISE,
                270,
            ],
        ];
    }

    /**
     * @dataProvider provideBoardData
     */
    public function testNewBoard__Ok(string $responseJson, array $board): void
    {
        $mockResponse = new MockResponse($responseJson);
        $mockClient = new MockHttpClient($mockResponse);

        $domainServiceClient = new DomainService(
            $mockClient,
            "http://domain-api"
        );

        $this->assertEquals($board, $domainServiceClient->newBoard(1));
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
        $domainServiceClient->newBoard(1);
    }

    /**
     * @dataProvider provideRotateRemainingTileData
     */
    public function testRotateRemainingTile__Ok(string $boardJson, array $board, Rotation $rotation, int $expectedRotation): void
    {
        $mockResponse = new MockResponse($boardJson);
        $mockClient = new MockHttpClient($mockResponse);

        $domainServiceClient = new DomainService(
            $mockClient,
            "http://domain-api"
        );

        $updatedBoard = $domainServiceClient->rotateRemainingTile($board, $rotation);
        $this->assertEquals($expectedRotation, $updatedBoard['remainingTile']['rotation']);
    }

    /**
     * @dataProvider provideRotateRemainingTileData
     */
    public function testRotateRemainingTile__InternalServererror(string $boardJson, array $board, Rotation $rotation, int $expectedRotation): void
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
        $domainServiceClient->rotateRemainingTile($board, $rotation);
    }
}
