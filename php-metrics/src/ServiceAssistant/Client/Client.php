<?php
declare(strict_types=1);

namespace Dimitriin\ServiceAssistant\Client;

use Dimitriin\ServiceAssistant\Protocol\Payload\Packet;
use Socket\Raw\Socket;

final class Client implements ClientInterface
{
    /**
     * @var Socket
     */
    private $socket;

    /**
     * Client constructor.
     *
     * @param Socket $socket
     */
    public function __construct(Socket $socket)
    {
        $this->socket = $socket;
    }

    public function send(Packet $packet): void
    {
        $this->socket->write($packet->serializeToString());
    }
}