/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.protobuf.CreateRoomReq', null, global);
goog.exportSymbol('proto.protobuf.PLAYER_STATE', null, global);
goog.exportSymbol('proto.protobuf.Player', null, global);
goog.exportSymbol('proto.protobuf.Room', null, global);
goog.exportSymbol('proto.protobuf.RoomState', null, global);
goog.exportSymbol('proto.protobuf.Round', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protobuf.Room = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protobuf.Room, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.protobuf.Room.displayName = 'proto.protobuf.Room';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protobuf.Room.prototype.toObject = function(opt_includeInstance) {
  return proto.protobuf.Room.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protobuf.Room} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.Room.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protobuf.Room}
 */
proto.protobuf.Room.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protobuf.Room;
  return proto.protobuf.Room.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protobuf.Room} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protobuf.Room}
 */
proto.protobuf.Room.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protobuf.Room.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protobuf.Room.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protobuf.Room} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.Room.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.protobuf.Room.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.protobuf.Room.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protobuf.CreateRoomReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protobuf.CreateRoomReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.protobuf.CreateRoomReq.displayName = 'proto.protobuf.CreateRoomReq';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protobuf.CreateRoomReq.prototype.toObject = function(opt_includeInstance) {
  return proto.protobuf.CreateRoomReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protobuf.CreateRoomReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.CreateRoomReq.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protobuf.CreateRoomReq}
 */
proto.protobuf.CreateRoomReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protobuf.CreateRoomReq;
  return proto.protobuf.CreateRoomReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protobuf.CreateRoomReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protobuf.CreateRoomReq}
 */
proto.protobuf.CreateRoomReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protobuf.CreateRoomReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protobuf.CreateRoomReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protobuf.CreateRoomReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.CreateRoomReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protobuf.RoomState = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.protobuf.RoomState.repeatedFields_, null);
};
goog.inherits(proto.protobuf.RoomState, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.protobuf.RoomState.displayName = 'proto.protobuf.RoomState';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.protobuf.RoomState.repeatedFields_ = [2,4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protobuf.RoomState.prototype.toObject = function(opt_includeInstance) {
  return proto.protobuf.RoomState.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protobuf.RoomState} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.RoomState.toObject = function(includeInstance, msg) {
  var f, obj = {
    room: (f = msg.getRoom()) && proto.protobuf.Room.toObject(includeInstance, f),
    playersList: jspb.Message.toObjectList(msg.getPlayersList(),
    proto.protobuf.Player.toObject, includeInstance),
    currentRound: (f = msg.getCurrentRound()) && proto.protobuf.Round.toObject(includeInstance, f),
    roundsList: jspb.Message.toObjectList(msg.getRoundsList(),
    proto.protobuf.Round.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protobuf.RoomState}
 */
proto.protobuf.RoomState.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protobuf.RoomState;
  return proto.protobuf.RoomState.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protobuf.RoomState} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protobuf.RoomState}
 */
proto.protobuf.RoomState.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.protobuf.Room;
      reader.readMessage(value,proto.protobuf.Room.deserializeBinaryFromReader);
      msg.setRoom(value);
      break;
    case 2:
      var value = new proto.protobuf.Player;
      reader.readMessage(value,proto.protobuf.Player.deserializeBinaryFromReader);
      msg.addPlayers(value);
      break;
    case 3:
      var value = new proto.protobuf.Round;
      reader.readMessage(value,proto.protobuf.Round.deserializeBinaryFromReader);
      msg.setCurrentRound(value);
      break;
    case 4:
      var value = new proto.protobuf.Round;
      reader.readMessage(value,proto.protobuf.Round.deserializeBinaryFromReader);
      msg.addRounds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protobuf.RoomState.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protobuf.RoomState.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protobuf.RoomState} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.RoomState.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRoom();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.protobuf.Room.serializeBinaryToWriter
    );
  }
  f = message.getPlayersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.protobuf.Player.serializeBinaryToWriter
    );
  }
  f = message.getCurrentRound();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.protobuf.Round.serializeBinaryToWriter
    );
  }
  f = message.getRoundsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      proto.protobuf.Round.serializeBinaryToWriter
    );
  }
};


/**
 * optional Room room = 1;
 * @return {?proto.protobuf.Room}
 */
proto.protobuf.RoomState.prototype.getRoom = function() {
  return /** @type{?proto.protobuf.Room} */ (
    jspb.Message.getWrapperField(this, proto.protobuf.Room, 1));
};


/** @param {?proto.protobuf.Room|undefined} value */
proto.protobuf.RoomState.prototype.setRoom = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.protobuf.RoomState.prototype.clearRoom = function() {
  this.setRoom(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protobuf.RoomState.prototype.hasRoom = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated Player players = 2;
 * @return {!Array.<!proto.protobuf.Player>}
 */
proto.protobuf.RoomState.prototype.getPlayersList = function() {
  return /** @type{!Array.<!proto.protobuf.Player>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.protobuf.Player, 2));
};


/** @param {!Array.<!proto.protobuf.Player>} value */
proto.protobuf.RoomState.prototype.setPlayersList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.protobuf.Player=} opt_value
 * @param {number=} opt_index
 * @return {!proto.protobuf.Player}
 */
proto.protobuf.RoomState.prototype.addPlayers = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.protobuf.Player, opt_index);
};


proto.protobuf.RoomState.prototype.clearPlayersList = function() {
  this.setPlayersList([]);
};


/**
 * optional Round current_round = 3;
 * @return {?proto.protobuf.Round}
 */
proto.protobuf.RoomState.prototype.getCurrentRound = function() {
  return /** @type{?proto.protobuf.Round} */ (
    jspb.Message.getWrapperField(this, proto.protobuf.Round, 3));
};


/** @param {?proto.protobuf.Round|undefined} value */
proto.protobuf.RoomState.prototype.setCurrentRound = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.protobuf.RoomState.prototype.clearCurrentRound = function() {
  this.setCurrentRound(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protobuf.RoomState.prototype.hasCurrentRound = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * repeated Round rounds = 4;
 * @return {!Array.<!proto.protobuf.Round>}
 */
proto.protobuf.RoomState.prototype.getRoundsList = function() {
  return /** @type{!Array.<!proto.protobuf.Round>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.protobuf.Round, 4));
};


/** @param {!Array.<!proto.protobuf.Round>} value */
proto.protobuf.RoomState.prototype.setRoundsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.protobuf.Round=} opt_value
 * @param {number=} opt_index
 * @return {!proto.protobuf.Round}
 */
proto.protobuf.RoomState.prototype.addRounds = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.protobuf.Round, opt_index);
};


proto.protobuf.RoomState.prototype.clearRoundsList = function() {
  this.setRoundsList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protobuf.Round = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.protobuf.Round.repeatedFields_, null);
};
goog.inherits(proto.protobuf.Round, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.protobuf.Round.displayName = 'proto.protobuf.Round';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.protobuf.Round.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protobuf.Round.prototype.toObject = function(opt_includeInstance) {
  return proto.protobuf.Round.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protobuf.Round} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.Round.toObject = function(includeInstance, msg) {
  var f, obj = {
    number: jspb.Message.getFieldWithDefault(msg, 1, 0),
    winner: (f = msg.getWinner()) && proto.protobuf.Player.toObject(includeInstance, f),
    playersList: jspb.Message.toObjectList(msg.getPlayersList(),
    proto.protobuf.Player.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protobuf.Round}
 */
proto.protobuf.Round.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protobuf.Round;
  return proto.protobuf.Round.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protobuf.Round} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protobuf.Round}
 */
proto.protobuf.Round.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setNumber(value);
      break;
    case 2:
      var value = new proto.protobuf.Player;
      reader.readMessage(value,proto.protobuf.Player.deserializeBinaryFromReader);
      msg.setWinner(value);
      break;
    case 3:
      var value = new proto.protobuf.Player;
      reader.readMessage(value,proto.protobuf.Player.deserializeBinaryFromReader);
      msg.addPlayers(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protobuf.Round.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protobuf.Round.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protobuf.Round} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.Round.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNumber();
  if (f !== 0) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = message.getWinner();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.protobuf.Player.serializeBinaryToWriter
    );
  }
  f = message.getPlayersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.protobuf.Player.serializeBinaryToWriter
    );
  }
};


/**
 * optional uint64 number = 1;
 * @return {number}
 */
proto.protobuf.Round.prototype.getNumber = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.protobuf.Round.prototype.setNumber = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional Player winner = 2;
 * @return {?proto.protobuf.Player}
 */
proto.protobuf.Round.prototype.getWinner = function() {
  return /** @type{?proto.protobuf.Player} */ (
    jspb.Message.getWrapperField(this, proto.protobuf.Player, 2));
};


/** @param {?proto.protobuf.Player|undefined} value */
proto.protobuf.Round.prototype.setWinner = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.protobuf.Round.prototype.clearWinner = function() {
  this.setWinner(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protobuf.Round.prototype.hasWinner = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * repeated Player players = 3;
 * @return {!Array.<!proto.protobuf.Player>}
 */
proto.protobuf.Round.prototype.getPlayersList = function() {
  return /** @type{!Array.<!proto.protobuf.Player>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.protobuf.Player, 3));
};


/** @param {!Array.<!proto.protobuf.Player>} value */
proto.protobuf.Round.prototype.setPlayersList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.protobuf.Player=} opt_value
 * @param {number=} opt_index
 * @return {!proto.protobuf.Player}
 */
proto.protobuf.Round.prototype.addPlayers = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.protobuf.Player, opt_index);
};


proto.protobuf.Round.prototype.clearPlayersList = function() {
  this.setPlayersList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protobuf.Player = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protobuf.Player, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.protobuf.Player.displayName = 'proto.protobuf.Player';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protobuf.Player.prototype.toObject = function(opt_includeInstance) {
  return proto.protobuf.Player.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protobuf.Player} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.Player.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    state: jspb.Message.getFieldWithDefault(msg, 3, 0),
    score: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protobuf.Player}
 */
proto.protobuf.Player.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protobuf.Player;
  return proto.protobuf.Player.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protobuf.Player} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protobuf.Player}
 */
proto.protobuf.Player.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {!proto.protobuf.PLAYER_STATE} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setScore(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protobuf.Player.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protobuf.Player.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protobuf.Player} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protobuf.Player.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getScore();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.protobuf.Player.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.protobuf.Player.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.protobuf.Player.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.protobuf.Player.prototype.setName = function(value) {
  jspb.Message.setField(this, 2, value);
};


/**
 * optional PLAYER_STATE state = 3;
 * @return {!proto.protobuf.PLAYER_STATE}
 */
proto.protobuf.Player.prototype.getState = function() {
  return /** @type {!proto.protobuf.PLAYER_STATE} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {!proto.protobuf.PLAYER_STATE} value */
proto.protobuf.Player.prototype.setState = function(value) {
  jspb.Message.setField(this, 3, value);
};


/**
 * optional int64 score = 4;
 * @return {number}
 */
proto.protobuf.Player.prototype.getScore = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.protobuf.Player.prototype.setScore = function(value) {
  jspb.Message.setField(this, 4, value);
};


/**
 * @enum {number}
 */
proto.protobuf.PLAYER_STATE = {
  UNKNOWN: 0,
  JOINED: 1,
  PLAYING: 2,
  DISCONNECTED: 3
};

goog.object.extend(exports, proto.protobuf);
