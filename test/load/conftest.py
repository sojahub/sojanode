import pytest
import sojatool_path
import sojatool.test_utils


@pytest.fixture(scope="function")
def ctx(request):
    yield from sojatool.test_utils.pytest_ctx_fixture(request)


@pytest.fixture(autouse=True)
def test_wrapper_fixture():
    sojatool.test_utils.pytest_test_wrapper_fixture()
