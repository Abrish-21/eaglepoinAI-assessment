function delay(ms) { // I can simply assign a delay but this makes the code cleaner
    return new Promise(resolve => setTimeout(resolve, ms));
}

// Mock API function that randomly succeeds or fails
function mockApiCall() {
    return new Promise((resolve, reject) => {
        const succeed = Math.random() > 0.8; // to see the retry effect so often 
        setTimeout(() => {
            if (succeed) {
                resolve({ data: "Success!" });
            } else {
                reject(new Error("API call failed"));
            }
        }, 300); // simulate network latency
    });
}

// Fetch with retry function
function fetchWithRetry(url, options = {}, retries = 3) {
    return new Promise((resolve, reject) => {
        mockApiCall(url, options) // using mockApiCall instead of fetch for demonstration
            .then(resolve)
            .catch(async (error) => {
                if (retries === 1) return reject(error);

                console.log(`Attempt failed. Retrying in 1 second... (${retries - 1} retries left)`);
                await delay(1000); // wait 1 second before retrying
                resolve(fetchWithRetry(url, options, retries - 1)); // The recursion here will continue until retries are exhausted
            });
    });
}



// Usage example
fetchWithRetry("https://example.com/data", {}, 5)
    .then(data => console.log("Fetched:", data))
    .catch(err => console.error("Failed:", err.message));
