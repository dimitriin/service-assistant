<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\ServiceAssistant;

use Dimitriin\Metrics\Core\CounterInterface;
use Dimitriin\ServiceAssistant\Client\ClientInterface;
use Dimitriin\ServiceAssistant\Protocol\Payload\CounterAddCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\CounterIncCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\Packet;

final class Counter implements CounterInterface
{
    /**
     * @var ClientInterface
     */
    private $client;

    /**
     * @var string
     */
    private $name;

    /**
     * Counter constructor.
     *
     * @param ClientInterface $client
     * @param string          $name
     */
    public function __construct(ClientInterface $client, string $name)
    {
        $this->client = $client;
        $this->name = $name;
    }

    public function inc(array $labels): void
    {
        $cmd = new CounterIncCmd();
        $cmd->setName($this->name);
        $cmd->setLabels($labels);

        $packet = new Packet();
        $packet->setCounterIncCmd($cmd);

        $this->client->send($packet);
    }

    public function incBy(int $val, array $labels): void
    {
        $cmd = new CounterAddCmd();
        $cmd->setName($this->name);
        $cmd->setLabels($labels);
        $cmd->setValue(floatval($val));

        $packet = new Packet();
        $packet->setCounterAddCmd($cmd);

        $this->client->send($packet);
    }
}