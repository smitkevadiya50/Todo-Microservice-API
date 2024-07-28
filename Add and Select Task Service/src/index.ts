import express, { ErrorRequestHandler } from 'express';
import { Pool } from 'pg';

const app = express();
const port = 3001;

const pool = new Pool({
    user: 'postgres',
    host: 'localhost',
    database: 'todoapp',
    password: '2002',
    port: 5432,
});

app.use(express.json());

app.post('/add', async (req, res) => {
    const { description } = req.body;
    try {
        const result = await pool.query('INSERT INTO tasks (description) VALUES ($1) RETURNING *', [description]);
        res.status(201).json(result.rows[0]);
    } catch (err) {
        res.status(500).json({ error: err });
    }
});

app.get('/', async (req, res) => {
    try {
        const result = await pool.query('SELECT * from tasks');
        if (result.rowCount === 0) {
            return res.status(404).json({ error: 'Task not found' });
        }
        res.status(200).json(result.rows);
    } catch (err) {
        res.status(500).json({ error: err });
    }
});


app.listen(port,() => {
    console.log(`Service One Add Task is listening at port: ${port}`);
});