<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\ServiceAssistant;

use Dimitriin\Metrics\Core\HistogramInterface;
use Dimitriin\ServiceAssistant\Client\ClientInterface;
use Dimitriin\ServiceAssistant\Protocol\Payload\HistogramObserveCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\Packet;

final class Histogram implements HistogramInterface
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

    public function observe(float $val, array $labels): void
    {
        $cmd = new HistogramObserveCmd();
        $cmd->setName($this->name);
        $cmd->setLabels($labels);
        $cmd->setValue(floatval($val));

        $packet = new Packet();
        $packet->setHistogramObserveCmd($cmd);

        $this->client->send($packet);
    }

}