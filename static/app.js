function loadPackSizes() {
    fetch('/api/packsizes')
        .then(resp => resp.json())
        .then(data => {
            const list = document.getElementById('packsizes');
            list.innerHTML = ''; // clear existing list

            if (!Array.isArray(data) || data.length === 0) {
                const li = document.createElement('li');
                li.textContent = '(no pack sizes configured)';
                li.classList.add('list-group-item', 'text-muted');
                list.appendChild(li);
                return;
            }

            data.forEach(size => {
                const li = document.createElement('li');
                li.textContent = size;
                li.classList.add('list-group-item');
                list.appendChild(li);
            });
        })
        .catch(err => {
            console.error('Error fetching pack sizes:', err);
            const list = document.getElementById('packsizes');
            list.innerHTML = '<li class="list-group-item text-danger">(error loading pack sizes)</li>';
        });
}

document.getElementById('calculate').addEventListener('click', () => {
    const qty = document.getElementById('quantity').value;
    const resultBox = document.getElementById('result');
    resultBox.classList.remove('bg-danger', 'text-white'); // reset error style

    if (!qty || qty <= 0) {
        alert('Please enter a positive quantity');
        return;
    }

    fetch('/api/calculate?quantity=' + qty)
        .then(resp => {
            if (!resp.ok) return resp.text().then(text => { throw new Error(text) });
            return resp.json();
        })
        .then(data => {
            let output = `Total Items: ${data.total_items}\n`;
            output += 'Packs:\n';
            data.packs.forEach(p => {
                output += `  ${p.count} x ${p.size}\n`;
            });
            resultBox.textContent = output;
        })
        .catch(err => {
            console.error('Error calculating packs:', err);
            resultBox.textContent = 'Error: ' + err.message;
            resultBox.classList.add('bg-danger', 'text-white');
        });
});

document.getElementById('addSize').addEventListener('click', () => {
    const newSize = parseInt(document.getElementById('newSize').value);
    if (!newSize || newSize <= 0) {
        alert('Please enter a positive number');
        return;
    }

    fetch('/api/packsizes/add', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ size: newSize })
    })
        .then(resp => {
            if (resp.ok) {
                loadPackSizes();
            } else {
                return resp.text().then(text => { throw new Error(text) });
            }
        })
        .catch(err => {
            console.error('Error adding pack size:', err);
        });
});

document.getElementById('removeSize').addEventListener('click', () => {
    const delSize = parseInt(document.getElementById('deleteSize').value);
    if (!delSize || delSize <= 0) {
        alert('Please enter a positive number');
        return;
    }

    fetch('/api/packsizes/delete?size=' + delSize, { method: 'DELETE' })
        .then(resp => {
            if (resp.ok) {
                loadPackSizes();
            } else {
                return resp.text().then(text => { throw new Error(text) });
            }
        })
        .catch(err => {
            console.error('Error removing pack size:', err);
        });
});

document.getElementById('removeAllSizes').addEventListener('click', () => {
    if (!confirm('Are you sure you want to remove ALL pack sizes?')) return;

    fetch('/api/packsizes/deleteall', { method: 'DELETE' })
        .then(resp => {
            if (resp.ok) {
                loadPackSizes();
            } else {
                return resp.text().then(text => { throw new Error(text) });
            }
        })
        .catch(err => {
            console.error('Error removing all pack sizes:', err);
        });
});

// Initial load
loadPackSizes();
