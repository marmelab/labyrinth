<?php

namespace App\Tests\Unit\Controller;

use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;
use Symfony\Component\HttpClient\MockHttpClient;
use Symfony\Component\HttpClient\Response\MockResponse;
use Symfony\Contracts\HttpClient\Exception\ServerExceptionInterface;

use App\Controller\HomeController;
use App\Service\DomainServiceInterface;

class HomeControllerTest extends WebTestCase
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
                        'treasure' => 'A'
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
                        'treasure' => 'B'
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
                        'treasure' => 'C'
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

    function testIndex()
    {
        $client = static::createClient();

        $container = static::getContainer();

        $domainServiceMock = $this->createMock(DomainServiceInterface::class);
        $domainServiceMock->expects(self::once())
            ->method('newBoard')
            ->willReturn(self::EXPECTED_BOARD);

        $container->set(DomainServiceInterface::class, $domainServiceMock);

        $crawler = $client->request('GET', '/');

        $this->assertResponseIsSuccessful();
        $this->assertCount(9, $crawler->filter('.tile'));
    }
}