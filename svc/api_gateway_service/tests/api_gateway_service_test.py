import json
from api_gateway_service.api import create_app
from unittest import TestCase

base_url = '/v1/links'


class APIGatewayServiceTest(TestCase):
    def setUp(self):
        self.app = create_app()

        self.test_app = self.app.test_client()

        # add links
        params = dict(username='Brijesh Lakkad',
                      url='https://tutsplus.com/authors/Brijesh Lakkad-sayfan',
                      title='Brijesh Lakkad Tuts+ articles')
        self.test_app.post(base_url, data=params)

    def tearDown(self):
        params = dict(username='Brijesh Lakkad',
                      url='https://tutsplus.com/authors/Brijesh Lakkad-sayfan')
        self.test_app.post(base_url, data=params)

    def test_get_links(self):
        url = base_url + '?username=Brijesh Lakkad'
        response = self.test_app.get(url, headers='')
        result = json.loads(response.data)
        expected = {}
        self.assertEqual(expected, result)
