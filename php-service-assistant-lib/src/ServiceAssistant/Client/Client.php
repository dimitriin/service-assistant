<?php
declare(strict_types=1);

namespace Dimitriin\ServiceAssistant\Client;

use Dimitriin\ServiceAssistant\Protocol\Payload\Packet;
use Psr\Log\LoggerAwareTrait;
use Psr\Log\NullLogger;
use Socket\Raw\Socket;
use Throwable;

final class Client implements ClientInterface
{
    use LoggerAwareTrait;

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
        $this->logger = new NullLogger();
    }

    public function send(Packet $packet): void
    {
        try {
            $this->socket->write($packet->serializeToString());
        } catch (Throwable $e) {
            $this->logger->error("Service assistant send packet exception",  [
                'exception' => $e->getMessage(),
                'packet' => $packet->serializeToString(),
            ]);
        }
    }
}