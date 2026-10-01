import http from 'node:http';

const PORT = 8086;

// Helper: simulate CPU-bound synchronous computation
function burnCpu(durationMs: number): number {
  const start = Date.now();
  let iterations = 0;
  while (Date.now() - start < durationMs) {
    iterations++;
    // Simple math computation in tight synchronous loop
    Math.sqrt(iterations * 1.5);
  }
  return iterations;
}

// Helper: simulate asynchronous I/O wait (database query, downstream RPC, disk read)
function asyncWait(durationMs: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, durationMs));
}

const server = http.createServer(async (req, res) => {
  const url = new URL(req.url || '/', `http://${req.headers.host}`);
  const startTime = Date.now();

  if (url.pathname === '/health') {
    res.writeHead(200, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({ status: 'ok' }));
    return;
  }

  if (url.pathname === '/cpu') {
    // Synchronously burns CPU for 150ms, starving all other concurrent requests on the single thread
    const iterations = burnCpu(150);
    const elapsed = Date.now() - startTime;

    res.writeHead(200, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({
      workload: 'cpu_bound',
      elapsed_ms: elapsed,
      cpu_burned: true,
      iterations,
    }));
    return;
  }

  if (url.pathname === '/io') {
    // Asynchronously pauses for 150ms; thread yields back to event loop, CPU usage remains ~0%
    await asyncWait(150);
    const elapsed = Date.now() - startTime;

    res.writeHead(200, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({
      workload: 'async_io_wait',
      elapsed_ms: elapsed,
      cpu_burned: false,
    }));
    return;
  }

  res.writeHead(404, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({ error: 'not_found' }));
});

server.listen(PORT, () => {
  console.log(`Async profiling server listening on http://localhost:${PORT}`);
  console.log(`Endpoints: /health, /cpu (CPU-bound), /io (Async waiting)`);
});
