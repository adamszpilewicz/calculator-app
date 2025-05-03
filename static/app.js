function loadPackSizes() {
    fetch('/api/packsizes')
        .then(resp => resp.json())
        .then(data => {
            const list = document.getElementById('packsizes');
            list.innerHTML = ''; // clear existing list
            data.forEach(size => {
                const li = document.createElement('li');
                li.textContent = size;
                list.appendChild(li);
            });
        })
        .catch(err => console.error('Error fetching pack sizes:', err));
}

document.getElementById('calculate').addEventListener('click', () => {
    const qty = document.getElementById('quantity').value;
    if (!qty || qty <= 0) {
        alert('Please enter a positive quantity');
        return;
    }

    fetch('/api/calculate?quantity=' + qty)
        .then(resp => {
            if (!resp.ok) {
                return resp.text().then(text => { throw new Error(text) });
            }
            return resp.json();
        })
        .then(data => {
            let output = `Total Items: ${data.total_items}\n`;
            output += 'Packs:\n';
            data.packs.forEach(p => {
                output += `  ${p.count} x ${p.size}\n`;
            });
            document.getElementById('result').textContent = output;
        })
        .catch(err => {
            console.error('Error calculating packs:', err);
            alert('Error: ' + err.message);
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
                alert('Pack size added');
                loadPackSizes();
            } else {
                return resp.text().then(text => { throw new Error(text) });
            }
        })
        .catch(err => {
            console.error('Error adding pack size:', err);
            alert('Error: ' + err.message);
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
                alert('Pack size removed');
                loadPackSizes();
            } else {
                return resp.text().then(text => { throw new Error(text) });
            }
        })
        .catch(err => {
            console.error('Error removing pack size:', err);
            alert('Error: ' + err.message);
        });
});

// Load initial pack sizes on page load
loadPackSizes();
