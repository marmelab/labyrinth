<?php

namespace App\Tests\Unit\Service;

use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;
use Symfony\Component\HttpClient\MockHttpClient;
use Symfony\Component\HttpClient\Response\MockResponse;
use Symfony\Contracts\HttpClient\Exception\ServerExceptionInterface;

use App\Service\DomainService;

class DomainServiceTest extends KernelTestCase
{
    private const EXPECTED_BOARD = [
        'tiles' => [
            [
                [
                    'tile' => [
                        'shape' => 2,
                        'treasure' => '·'
                    ],
                    'rotation' => 270
                ],
                [
                    'tile' => [
                        'shape' => 1,
                        'treasure' => 'H'
                    ],
                    'rotation' => 270
                ],
                [
                    'tile' => [
                        'shape' => 2,
                        'treasure' => '·'
                    ],
                    'rotation' => 270
                ],
            ],
            [
                [
                    'tile' => [
                        'shape' => 2,
                        'treasure' => '·'
                    ],
                    'rotation' => 90
                ],
                [
                    'tile' => [
                        'shape' => 1,
                        'treasure' => 'F'
                    ],
                    'rotation' => 0
                ],
                [
                    'tile' => [
                        'shape' => 2,
                        'treasure' => '·'
                    ],
                    'rotation' => 270
                ],
            ],
        ],
        'remainingTile' => [
            'tile' => [
                'shape' => 0,
                'treasure' => '·'
            ],
            'rotation' => 0
        ],
        'players' => [
            [
                'color' => 0,
                'position' => [
                    'line' => 0,
                    'row' => 0
                ],
                'targets' => [
                    'C',
                    'M',
                    'J',
                    'Q',
                ],
                'score' => 0
            ]
        ],
        'remainingPlayers' => [
            0
        ],
        'currentPlayerIndex' => 0,
        'gameState' => 0
    ];

    public function testNewBoard__Ok()
    {
        $mockResponse = new MockResponse(json_encode(self::EXPECTED_BOARD));
        $mockClient = new MockHttpClient($mockResponse);

        $domainServiceClient = new DomainService(
            $mockClient,
            "http://domain-api"
        );

        $this->assertEquals(self::EXPECTED_BOARD, $domainServiceClient->newBoard());
    }

    public function testNewBoard__InternalServererror()
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