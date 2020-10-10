<?php
declare(strict_types=1);

namespace Dimitriin\Metrics\ServiceAssistant;

use Dimitriin\Metrics\Core\CounterInterface;
use Dimitriin\Metrics\Core\Exception\MetricAlreadyRegisteredException;
use Dimitriin\Metrics\Core\Exception\MetricNotRegisteredException;
use Dimitriin\Metrics\Core\GaugeInterface;
use Dimitriin\Metrics\Core\HistogramInterface;
use Dimitriin\Metrics\Core\RegistryInterface;
use Dimitriin\ServiceAssistant\Client\ClientInterface;
use Dimitriin\ServiceAssistant\Protocol\Payload\CounterRegisterCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\GaugeRegisterCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\HistogramRegisterCmd;
use Dimitriin\ServiceAssistant\Protocol\Payload\Packet;

final class Registry implements RegistryInterface
{
    /**
     * @var ClientInterface
     */
    private $client;

    private $metrics;

    /**
     * Registry constructor.
     *
     * @param ClientInterface $client
     */
    public function __construct(ClientInterface $client)
    {
        $this->client = $client;
    }

    public function registerCounter(string $name, string $help, array $labels): void
    {
        $this->assertMetricNotRegistered($name);

        $cmd = new CounterRegisterCmd();
        $cmd->setName($name);
        $cmd->setHelp($help);
        $cmd->setLabels($labels);

        $packet = new Packet();
        $packet->setCounterRegisterCmd($cmd);
        
        $this->client->send($packet);

        $this->metrics[$name] = new Counter($this->client, $name);
    }

    public function registerGauge(string $name, string $help, array $labels): void
    {
        $this->assertMetricNotRegistered($name);

        $cmd = new GaugeRegisterCmd();
        $cmd->setName($name);
        $cmd->setHelp($help);
        $cmd->setLabels($labels);

        $packet = new Packet();
        $packet->setGaugeRegisterCmd($cmd);

        $this->client->send($packet);

        $this->metrics[$name] = new Gauge($this->client, $name);
    }

    public function registerHistogram(string $name, string $help, array $labels, ?array $buckets = null): void
    {
        $this->assertMetricNotRegistered($name);

        $cmd = new HistogramRegisterCmd();
        $cmd->setName($name);
        $cmd->setHelp($help);
        $cmd->setLabels($labels);
        $cmd->setBuckets($buckets ?? []);

        $packet = new Packet();
        $packet->setHistogramRegisterCmd($cmd);

        $this->client->send($packet);

        $this->metrics[$name] = new Histogram($this->client, $name);
    }

    public function getCounter(string $name): CounterInterface
    {
        $this->assertMetricRegistered($name);

        return $this->metrics[$name];
    }

    public function getGauge(string $name): GaugeInterface
    {
        $this->assertMetricRegistered($name);

        return $this->metrics[$name];
    }

    public function getHistogram(string $name): HistogramInterface
    {
        $this->assertMetricRegistered($name);

        return $this->metrics[$name];
    }

    private function assertMetricNotRegistered(string $name): void {
        if(isset($this->metrics[$name])) {
            throw new MetricAlreadyRegisteredException($name);
        }
    }

    private function assertMetricRegistered(string $name): void {
        if(!isset($this->metrics[$name])) {
            throw new MetricNotRegisteredException($name);
        }
    }
}