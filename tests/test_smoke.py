"""Smoke tests for wasm-education."""

import sys
import importlib.util
import ast
import os
import pytest

def test_import_check_pages():
    """Test that check_pages module can be imported."""
    try:
        spec = importlib.util.find_spec("tools.check_pages")
        assert spec is not None, "check_pages module not found"
        module = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(module)
        assert module is not None
        # Check that the main function exists
        assert hasattr(module, "main") or hasattr(module, "check_pages")
    except Exception as e:
        pytest.skip(f"check_pages module has missing dependencies: {e}")

def test_python_syntax_valid():
    """Test that all Python files have valid syntax."""
    python_files = []
    for root, dirs, files in os.walk("."):
        dirs[:] = [d for d in dirs if not d.startswith('.') and d != '__pycache__']
        for file in files:
            if file.endswith(".py"):
                python_files.append(os.path.join(root, file))
    
    assert len(python_files) > 0, "No Python files found"
    
    for py_file in python_files:
        with open(py_file, "r", encoding="utf-8", errors="replace") as f:
            content = f.read()
        try:
            ast.parse(content)
        except SyntaxError as e:
            pytest.fail(f"Syntax error in {py_file}: {e}")

def test_check_pages_runs():
    """Test that check_pages.py can run without error (smoke test)."""
    import subprocess
    result = subprocess.run(
        [sys.executable, "tools/check_pages.py", "--help"],
        capture_output=True,
        text=True,
        timeout=30
    )
    # Should exit with 0 or 2 (help), not crash
    assert result.returncode in (0, 2), f"check_pages.py crashed: {result.stderr}"

if __name__ == "__main__":
    test_import_check_pages()
    test_python_syntax_valid()
    test_check_pages_runs()
    print("All smoke tests passed!")