<?php

namespace App\Service;

use Symfony\Contracts\HttpClient\HttpClientInterface;

class DomainService implements DomainServiceInterface
{

    public function __construct(
        private HttpClientInterface $httpClient,
        private string $domainServiceUrl
    )
    {

    }

    public function newBoard(): array
    {
        $response = $this->httpClient->request("POST", "{$this->domainServiceUrl}/new");
        return $response->toArray();
    }


    function rotateRemainingTile(array $board, Rotation $rotation): array
    {
        $response = $this->httpClient->request("POST", "{$this->domainServiceUrl}/rotate-remaining", [
            'headers' => [
                'Content-Type' => 'application/json',
            ],
            'body' => json_encode([
                'board' => $board,
                'rotation' => $rotation,
            ]),
        ]);

        return $response->toArray();
    }
}