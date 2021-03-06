<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: payload.proto

namespace Dimitriin\ServiceAssistant\Protocol\Payload;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>Dimitriin.ServiceAssistant.Protocol.Payload.ReadyBit</code>
 */
class ReadyBit extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>uint64 ttl = 1;</code>
     */
    private $ttl = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $ttl
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Payload::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>uint64 ttl = 1;</code>
     * @return int|string
     */
    public function getTtl()
    {
        return $this->ttl;
    }

    /**
     * Generated from protobuf field <code>uint64 ttl = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setTtl($var)
    {
        GPBUtil::checkUint64($var);
        $this->ttl = $var;

        return $this;
    }

}

