<?php
declare(strict_types=1);

namespace Dimitriin\ServiceAssistant\Client;

use Dimitriin\ServiceAssistant\Protocol\Payload\Packet;

interface ClientInterface
{
    public function send(Packet $packet): void;
}