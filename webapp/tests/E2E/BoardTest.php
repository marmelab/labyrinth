<?php

namespace App\Tests\E2E;

use Symfony\Component\Panther\PantherTestCase;

use Facebook\WebDriver\WebDriverDimension;

class BoardTest extends PantherTestCase
{
    public function testRotateRemainingTileClockwise(): void
    {
        $client = static::createPantherClient([
            'browser' => PantherTestCase::FIREFOX,
        ]);
        $client->request('GET', $_ENV['WEBAPP_URL']);
        $client->manage()->window()->setSize(new WebDriverDimension(1920, 1080));

        $crawler = $client->waitFor('.tile--remaining');

        $remainingTileFilter = $crawler->filter('.tile--remaining');
        $this->assertCount(1, $remainingTileFilter);
        $initialClasses = $remainingTileFilter->first()->getAttribute("class");

        $client->submitForm('submit-rotate-remaining-clockwise');

        $crawler = $client->waitFor('.tile--remaining');

        $remainingTileFilter = $crawler->filter('.tile--remaining');
        $this->assertCount(1, $remainingTileFilter);

        $updatedClasses = $remainingTileFilter->first()->getAttribute("class");

        $this->assertNotEquals($initialClasses, $updatedClasses);
    }
    public function testRotateRemainingTileAntilockwise(): void
    {
        $client = static::createPantherClient([
            'browser' => PantherTestCase::FIREFOX,
        ]);
        $client->request('GET', $_ENV['WEBAPP_URL']);
        $client->manage()->window()->setSize(new WebDriverDimension(1920, 1080));

        $crawler = $client->waitFor('.tile--remaining');

        $remainingTileFilter = $crawler->filter('.tile--remaining');
        $this->assertCount(1, $remainingTileFilter);
        $initialClasses = $remainingTileFilter->first()->getAttribute("class");

        $client->submitForm('submit-rotate-remaining-anticlockwise');

        $crawler = $client->waitFor('.tile--remaining');

        $remainingTileFilter = $crawler->filter('.tile--remaining');
        $this->assertCount(1, $remainingTileFilter);

        $updatedClasses = $remainingTileFilter->first()->getAttribute("class");

        $this->assertNotEquals($initialClasses, $updatedClasses);
    }
}