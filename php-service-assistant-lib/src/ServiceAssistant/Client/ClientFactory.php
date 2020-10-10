<?php
declare(strict_types=1);

namespace Dimitriin\ServiceAssistant\Client;

use Socket\Raw\Factory;

final class ClientFactory
{
    public function __invoke(string $address) {
        $socketFactory = new Factory();
        $socket = $socketFactory->createClient($address);

        return new Client($socket);
    }
}