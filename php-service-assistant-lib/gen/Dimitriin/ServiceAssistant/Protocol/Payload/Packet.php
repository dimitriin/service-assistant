<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: payload.proto

namespace Dimitriin\ServiceAssistant\Protocol\Payload;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>Dimitriin.ServiceAssistant.Protocol.Payload.Packet</code>
 */
class Packet extends \Google\Protobuf\Internal\Message
{
    protected $Payload;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\ReadyBit $readyBit
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\HealthBit $healthBit
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\CounterRegisterCmd $counterRegisterCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\CounterIncCmd $counterIncCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\CounterAddCmd $counterAddCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\HistogramRegisterCmd $histogramRegisterCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\HistogramObserveCmd $histogramObserveCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeRegisterCmd $gaugeRegisterCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeIncCmd $gaugeIncCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeDecCmd $gaugeDecCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetCmd $gaugeSetCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeAddCmd $gaugeAddCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSubCmd $gaugeSubCmd
     *     @type \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetToCurrentTimeCmd $gaugeSetToCurrentTimeCmd
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Payload::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.ReadyBit readyBit = 1;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\ReadyBit
     */
    public function getReadyBit()
    {
        return $this->readOneof(1);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.ReadyBit readyBit = 1;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\ReadyBit $var
     * @return $this
     */
    public function setReadyBit($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\ReadyBit::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.HealthBit healthBit = 2;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\HealthBit
     */
    public function getHealthBit()
    {
        return $this->readOneof(2);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.HealthBit healthBit = 2;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\HealthBit $var
     * @return $this
     */
    public function setHealthBit($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\HealthBit::class);
        $this->writeOneof(2, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.CounterRegisterCmd counterRegisterCmd = 3;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\CounterRegisterCmd
     */
    public function getCounterRegisterCmd()
    {
        return $this->readOneof(3);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.CounterRegisterCmd counterRegisterCmd = 3;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\CounterRegisterCmd $var
     * @return $this
     */
    public function setCounterRegisterCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\CounterRegisterCmd::class);
        $this->writeOneof(3, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.CounterIncCmd counterIncCmd = 4;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\CounterIncCmd
     */
    public function getCounterIncCmd()
    {
        return $this->readOneof(4);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.CounterIncCmd counterIncCmd = 4;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\CounterIncCmd $var
     * @return $this
     */
    public function setCounterIncCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\CounterIncCmd::class);
        $this->writeOneof(4, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.CounterAddCmd counterAddCmd = 5;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\CounterAddCmd
     */
    public function getCounterAddCmd()
    {
        return $this->readOneof(5);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.CounterAddCmd counterAddCmd = 5;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\CounterAddCmd $var
     * @return $this
     */
    public function setCounterAddCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\CounterAddCmd::class);
        $this->writeOneof(5, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.HistogramRegisterCmd histogramRegisterCmd = 6;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\HistogramRegisterCmd
     */
    public function getHistogramRegisterCmd()
    {
        return $this->readOneof(6);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.HistogramRegisterCmd histogramRegisterCmd = 6;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\HistogramRegisterCmd $var
     * @return $this
     */
    public function setHistogramRegisterCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\HistogramRegisterCmd::class);
        $this->writeOneof(6, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.HistogramObserveCmd histogramObserveCmd = 7;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\HistogramObserveCmd
     */
    public function getHistogramObserveCmd()
    {
        return $this->readOneof(7);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.HistogramObserveCmd histogramObserveCmd = 7;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\HistogramObserveCmd $var
     * @return $this
     */
    public function setHistogramObserveCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\HistogramObserveCmd::class);
        $this->writeOneof(7, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeRegisterCmd gaugeRegisterCmd = 8;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeRegisterCmd
     */
    public function getGaugeRegisterCmd()
    {
        return $this->readOneof(8);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeRegisterCmd gaugeRegisterCmd = 8;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeRegisterCmd $var
     * @return $this
     */
    public function setGaugeRegisterCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeRegisterCmd::class);
        $this->writeOneof(8, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeIncCmd gaugeIncCmd = 9;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeIncCmd
     */
    public function getGaugeIncCmd()
    {
        return $this->readOneof(9);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeIncCmd gaugeIncCmd = 9;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeIncCmd $var
     * @return $this
     */
    public function setGaugeIncCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeIncCmd::class);
        $this->writeOneof(9, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeDecCmd gaugeDecCmd = 10;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeDecCmd
     */
    public function getGaugeDecCmd()
    {
        return $this->readOneof(10);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeDecCmd gaugeDecCmd = 10;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeDecCmd $var
     * @return $this
     */
    public function setGaugeDecCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeDecCmd::class);
        $this->writeOneof(10, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeSetCmd gaugeSetCmd = 11;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetCmd
     */
    public function getGaugeSetCmd()
    {
        return $this->readOneof(11);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeSetCmd gaugeSetCmd = 11;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetCmd $var
     * @return $this
     */
    public function setGaugeSetCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetCmd::class);
        $this->writeOneof(11, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeAddCmd gaugeAddCmd = 12;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeAddCmd
     */
    public function getGaugeAddCmd()
    {
        return $this->readOneof(12);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeAddCmd gaugeAddCmd = 12;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeAddCmd $var
     * @return $this
     */
    public function setGaugeAddCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeAddCmd::class);
        $this->writeOneof(12, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeSubCmd gaugeSubCmd = 13;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSubCmd
     */
    public function getGaugeSubCmd()
    {
        return $this->readOneof(13);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeSubCmd gaugeSubCmd = 13;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSubCmd $var
     * @return $this
     */
    public function setGaugeSubCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSubCmd::class);
        $this->writeOneof(13, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeSetToCurrentTimeCmd gaugeSetToCurrentTimeCmd = 14;</code>
     * @return \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetToCurrentTimeCmd
     */
    public function getGaugeSetToCurrentTimeCmd()
    {
        return $this->readOneof(14);
    }

    /**
     * Generated from protobuf field <code>.Dimitriin.ServiceAssistant.Protocol.Payload.GaugeSetToCurrentTimeCmd gaugeSetToCurrentTimeCmd = 14;</code>
     * @param \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetToCurrentTimeCmd $var
     * @return $this
     */
    public function setGaugeSetToCurrentTimeCmd($var)
    {
        GPBUtil::checkMessage($var, \Dimitriin\ServiceAssistant\Protocol\Payload\GaugeSetToCurrentTimeCmd::class);
        $this->writeOneof(14, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getPayload()
    {
        return $this->whichOneof("Payload");
    }

}

