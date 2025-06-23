from flask import Flask, render_template, request, jsonify, send_from_directory
import subprocess
import tempfile
import os
import shutil

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/run', methods=['POST'])
def run_code():
    code = request.json.get('code', '')
    if not code:
        return jsonify({'output': 'No code provided.'})
    # Write code to a temporary file
    with tempfile.NamedTemporaryFile(delete=False, suffix='.om', mode='w', encoding='utf-8') as tmp:
        tmp.write(code)
        tmp_path = tmp.name
    try:
        # Run the Go interpreter (main.exe) with the temp file as argument
        result = subprocess.run([
            os.path.join(os.path.dirname(__file__), '..', 'main.exe'), tmp_path
        ], capture_output=True, text=True, timeout=10)
        output = result.stdout + (result.stderr if result.stderr else '')
    except Exception as e:
        output = f'Error: {e}'
    finally:
        os.unlink(tmp_path)
    return jsonify({'output': output})

@app.route('/learn')
def learn():
    return render_template('learn.html')

@app.route('/tutorial')
def tutorial():
    return render_template('tutorial.html')

# Ensure omexamples are available in static/examples
examples_src = os.path.join(os.path.dirname(__file__), '..', 'omexamples')
examples_dst = os.path.join(os.path.dirname(__file__), 'static', 'examples')
if not os.path.exists(examples_dst):
    os.makedirs(examples_dst)
for fname in os.listdir(examples_src):
    shutil.copy(os.path.join(examples_src, fname), os.path.join(examples_dst, fname))

@app.route('/static/examples/<path:filename>')
def example_files(filename):
    return send_from_directory(os.path.join(app.root_path, 'static', 'examples'), filename)

@app.route('/example/<filename>')
def example_repl(filename):
    example_path = os.path.join(app.root_path, 'static', 'examples', filename)
    code = ''
    if os.path.exists(example_path):
        with open(example_path, 'r', encoding='utf-8') as f:
            code = f.read()
    return render_template('index.html', example_code=code)

@app.route('/download/om.exe')
def download_compiler():
    exe_path = os.path.join(os.path.dirname(__file__), '..', 'main.exe')
    return send_from_directory(os.path.dirname(exe_path), os.path.basename(exe_path), as_attachment=True)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0') 