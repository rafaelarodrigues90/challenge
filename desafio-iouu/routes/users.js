const UserControl = require('../controllers/UserControl');
const express = require('express');
const router = express.Router();

// Middleware
router.use(express.urlencoded({ extended: true }));
router.use(express.json());

// CRUD
router.post('/create', UserControl.create);
router.put('/update/:_id', UserControl.update);
router.get('/retrieve', UserControl.retrieve);
router.delete('/delete/:_id', UserControl.delete);


module.exports = router;
