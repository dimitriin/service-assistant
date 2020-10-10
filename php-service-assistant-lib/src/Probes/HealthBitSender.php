<?php
declare(strict_types=1);

namespace Dimitriin\Probes;

use Dimitriin\ServiceAssistant\Client\ClientInterface;
use Dimitriin\ServiceAssistant\Protocol\Payload\HealthBit;
use Dimitriin\ServiceAssistant\Protocol\Payload\Packet;

final class HealthBitSender implements HealthBitSenderInterface
{
    /**
     * @var ClientInterface
     */
    private $client;

    /**
     * @var int
     */
    private $ttl;

    /**
     * HealthBitSender constructor.
     *
     * @param ClientInterface $client
     * @param int|null        $ttl
     */
    public function __construct(ClientInterface $client, ?int $ttl = null)
    {
        $this->client = $client;
        $this->ttl = $ttl ?? 60;
    }

    public function send(?int $ttl = null): void
    {
        $bit = new HealthBit();
        $bit->setTtl($ttl ?? $this->ttl);

        $packet = new Packet();
        $packet->setHealthBit($bit);

        $this->client->send($packet);
    }
}