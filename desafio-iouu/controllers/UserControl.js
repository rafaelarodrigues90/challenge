const UserModel = require('../models/UserModel');

module.exports = {
  create: (req, res) => {
    const user = new UserModel({
        name: req.body.name,
        surname: req.body.surname,
        cpf: req.body.cpf,
        balance: req.body.balance,
        email: req.body.email
    })
    user.save()
      .then(result => {
        res.json({ success: true, result: result });
      })
      .catch(err => {
        res.json({ success: false, result: err });
      });
    },

    // BUG: EXIBIR USUÁRIO EDITADO
  update: (req, res) => {
    UserModel.update({_id: req.params._id}, req.params)
        .then(user => {
        if (!user) res.json({ success: false, result: "User does not exist" }); 
        res.json(user);
        })
        .catch(err => {
          res.json({ success: false, result: err });
        })
  },
  retrieve: (req, res) => {
    UserModel.find()
      .then(result => {
        if (!result) res.json({ success: false, result: "No results found" });

        res.json({ success: true, result: result });
      })
      .catch(err => res.json({success: false, result: err}));
  },
  delete: (req, res) => {
    UserModel.remove({_id: req.params._id})
      .then(result => {
        if (!result) res.json({ success: false, result: "No user was found with the ID" });
        res.json({ success: true, result: "User deleted" });
      })
      .catch(err => res.json({ success: false, result: err }));
  }
}