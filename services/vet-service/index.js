// index.js
const express = require("express");
const mongoose = require("mongoose");
const bodyParser = require("body-parser");
const { v4: uuidv4 } = require("uuid");

const app = express();
const port = 3000;

// Middleware
app.use(bodyParser.json());

// MongoDB connection
mongoose
  .connect("mongodb://localhost/vet-service", {
    useNewUrlParser: true,
    useUnifiedTopology: true,
  })
  .then(() => console.log("Connected to MongoDB"))
  .catch((err) => console.log("MongoDB connection error: ", err));

// Vet model (MongoDB Schema)
const vetSchema = new mongoose.Schema({
  id: { type: String, default: uuidv4 },
  name: String,
  specialty: String,
  yearsOfExperience: Number,
});

const Vet = mongoose.model("Vet", vetSchema);

// Routes
// Get all vets
app.get("/vets", async (req, res) => {
  try {
    const vets = await Vet.find();
    res.json(vets);
  } catch (err) {
    res.status(500).send("Error retrieving vets");
  }
});

// Get a specific vet by ID
app.get("/vets/:id", async (req, res) => {
  try {
    const vet = await Vet.findOne({ id: req.params.id });
    if (!vet) return res.status(404).send("Vet not found");
    res.json(vet);
  } catch (err) {
    res.status(500).send("Error retrieving vet");
  }
});

// Create a new vet
app.post("/vets", async (req, res) => {
  try {
    const newVet = new Vet(req.body);
    await newVet.save();
    res.status(201).json(newVet);
  } catch (err) {
    res.status(400).send("Error creating vet");
  }
});

// Update vet details
app.put("/vets/:id", async (req, res) => {
  try {
    const updatedVet = await Vet.findOneAndUpdate(
      { id: req.params.id },
      req.body,
      { new: true }
    );
    if (!updatedVet) return res.status(404).send("Vet not found");
    res.json(updatedVet);
  } catch (err) {
    res.status(400).send("Error updating vet");
  }
});

// Delete a vet
app.delete("/vets/:id", async (req, res) => {
  try {
    const deletedVet = await Vet.findOneAndDelete({ id: req.params.id });
    if (!deletedVet) return res.status(404).send("Vet not found");
    res.status(204).send();
  } catch (err) {
    res.status(500).send("Error deleting vet");
  }
});

// Start the server
app.listen(port, () => {
  console.log(`Vet Service running on http://localhost:${port}`);
});
