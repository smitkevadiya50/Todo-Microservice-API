import express from 'express';
import { Pool } from 'pg';

const app = express();
const port = 3003;

const pool = new Pool({
    user: 'postgres',
    host: 'localhost',
    database: 'todoapp',
    password: '2002',
    port: 5432,
});

app.use(express.json());

app.put('/complete/:id', async (req, res) => {
    const { id } = req.params;
    try {
        const result = await pool.query('UPDATE tasks SET completed = TRUE WHERE id = $1 RETURNING *', [id]);
        if (result.rowCount === 0) {
            return res.status(404).json({ error: 'Task not found' });
        }
        res.status(200).json(result.rows[0]);
    } catch (err) {
        res.status(500).json({ error: err });
    }
});

app.listen(port,() => {
    console.log(`Service Three Complete Task is listening at port: ${port}`);
});