"""example python api client for zlint"""
import hashlib
from typing import Dict


import requests


def lint(certificate: str, encoding: str) -> Dict:
    """
    Sends a certificate to a ZLint API endpoint and retrieves linting results.

    This function posts the given certificate to a specified ZLint API endpoint,
    depending on the encoding format. It processes the response to include additional
    information about the linting results, such as a count and hash of failed lints.

    Args:
        certificate (str): The certificate to be linted.
        encoding (str): The encoding format of the certificate ('der' or 'pem').

    Returns:
        Dict: A dictionary containing the lint results, including additional fields
              for failed lints, their count, and a hash. Returns None if an error occurs.
    """
    data = requests.post(
        url=f'http://localhost:8080/{encoding}',
        json={'certificate': certificate},
        timeout=2
    ).json()

    if data.get('error'):
        return None

    data['failed_lints'] = [
        lint
        for lint, data in data['lints'].items()
        if data['result'] not in ('pass', 'NA', 'NE', 'reserved', '')
    ]
    data['failed_lints_hash'] = hashlib.sha256("".join(data['failed_lints']).encode()).hexdigest()
    data['failed_lints_count'] = len(data['failed_lints'])
    data['lints'] = {lint: data['result'] for lint, data in data['lints'].items()}

    return data


def der_lint(certificate: str) -> Dict:
    """
    Processes a certificate encoded in DER format through lint.

    This is a convenience function that calls the `lint` function with 'der' as the
    encoding parameter.

    Args:
        certificate (str): The DER-encoded certificate to be linted.

    Returns:
        Dict: The result of the lint operation.
    """
    return lint(certificate=certificate, encoding='der')


def pem_lint(certificate: str) -> Dict:
    """
    Processes a certificate encoded in PEM format through lint.

    This is a convenience function that calls the `lint` function with 'pem' as the
    encoding parameter.

    Args:
        certificate (str): The PEM-encoded certificate to be linted.

    Returns:
        Dict: The result of the lint operation.
    """
    return lint(certificate=certificate, encoding='pem')
