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
}