<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\ServiceAssistant;

use Dimitriin\Metrics\Core\GaugeInterface;
use Dimitriin\ServiceAssistant\Client\ClientInterface;
use Dimitriin\ServiceAssistant\Protocol\Payload\GaugeAddCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\GaugeDecCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\GaugeIncCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSubCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\Packet;

final class Gauge implements GaugeInterface
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
        $cmd = new GaugeIncCmd();
        $cmd->setName($this->name);
        $cmd->setLabels($labels);

        $packet = new Packet();
        $packet->setGaugeIncCmd($cmd);

        $this->client->send($packet);
    }

    public function set(float $val, array $labels): void
    {
        $cmd = new GaugeSetCmd();
        $cmd->setName($this->name);
        $cmd->setLabels($labels);
        $cmd->setValue($val);

        $packet = new Packet();
        $packet->setGaugeSetCmd($cmd);

        $this->client->send($packet);
    }

    public function incBy(float $val, array $labels): void
    {
        $cmd = new GaugeAddCmd();
        $cmd->setName($this->name);
        $cmd->setLabels($labels);
        $cmd->setValue($val);

        $packet = new Packet();
        $packet->setGaugeAddCmd($cmd);

        $this->client->send($packet);
    }

    public function dec(array $labels): void
    {
        $cmd = new GaugeDecCmd();
        $cmd->setName($this->name);
        $cmd->setLabels($labels);

        $packet = new Packet();
        $packet->setGaugeDecCmd($cmd);

        $this->client->send($packet);
    }

    public function decBy(float $val, array $labels): void
    {
        $cmd = new GaugeSubCmd();
        $cmd->setName($this->name);
        $cmd->setLabels($labels);
        $cmd->setValue($val);

        $packet = new Packet();
        $packet->setGaugeSubCmd($cmd);

        $this->client->send($packet);
    }

}