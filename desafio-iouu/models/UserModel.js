const mongoose = require('mongoose');
const Schema = mongoose.Schema;
const Float = require('mongoose-float').loadType(mongoose, 2);

const UserSchema = new Schema({
    // _id: { type: mongoose.Schema.ObjectId },
    name: { type: String, required: true, max: 100},
    surname: { type: String, required: true, max: 100},
    cpf: { type: Number, required: true, unique: true, maxlength: 11, minlength: 11 },
    balance: { type: Float, require: true, min: 0.00 },
    email: { type: String, required: true }
});

module.exports = mongoose.model('user', UserSchema);