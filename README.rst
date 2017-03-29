Go Tokens Library
=================

.. image:: https://travis-ci.org/zalando-incubator/go-tokens.svg?branch=master
    :target: https://travis-ci.org/zalando-incubator/go-tokens

.. image:: https://codecov.io/github/zalando-incubator/go-tokens/coverage.svg?branch=master
    :target: https://codecov.io/github/zalando-incubator/go-tokens?branch=master

.. image:: https://goreportcard.com/badge/github.com/zalando-incubator/go-tokens
    :target: https://goreportcard.com/report/github.com/zalando-incubator/go-tokens

.. image:: https://godoc.org/github.com/zalando-incubator/go-tokens?status.svg
    :target: https://godoc.org/github.com-incubator/zalando/go-tokens


**go-tokens** is a library that refreshes your OAuth tokens before they expire. Simply provide go-tokens with your OAuth2 token endpoint and tokens and scopes that you want it to manage, and it will ensure that your managed tokens remain valid. It can use any custom credentials provider, and you can use your own credential providers.

Credentials Provider
--------------------

go-tokens currently contains implementations of the ``user.CredentialsProvider`` and the ``client.CredentialsProvider``
that you can use out of the box. The simplest providers just return some static values used at the time of creation. More complex providers fetch credentials from JSON files (user.json and client.json) from a folder defined in the ``CREDENTIALS_DIR`` environment variable.

User Credentials
~~~~~~~~~~~~~~~~

User credentials consist, quite simply, of a username and a password. Any type that implements the ``user.Credentials`` should be able to provide them. Implement the ``user.CredentialsProvider`` interface for any custom type that is able to provide ``user.Credentials``.

For a simple example, check the `user/static.go`_ file.

Client Credentials
~~~~~~~~~~~~~~~~~~

Client credentials consist of a client ID and client secret, and are very similar to user credentials. Any type that implements the ``client.Credentials`` should be able to provide them. Implement the ``client.CredentialsProvider`` interface for any custom type that can provide ``client.Credentials``.

For a simple example, check the `client/static.go`_ file.

User Guide
----------

.. code-block:: go

    url := "https://example.com/oauth2/access_token"

    // You can manage multiple tokens with different scopes
    reqs := []tokens.ManagementRequest{
        tokens.NewPasswordRequest("test1", "foo.read"),
        tokens.NewPasswordRequest("test2", "user.email", "user.name"),
    }

    // Manager creation tries to obtain all tokens synchronously initially
    tokensManager, err := tokens.Manage(url, reqs)
    if err != nil {
        log.Fatal(err)
    }

    if test1, err := tokensManager.Get("test1"); err == nil {
        // Do something with the access token "test1"
    }

    if test2, err := tokensManager.Get("test2"); err == nil {
        // Do something with the access token "test2"
    }

This example would create a token manager using the JSON files credentials providers and a refresh threshold of 60% of the token validity time.

Manager Options
~~~~~~~~~~~~~~~

You can customize the behavior of the Manager with the following options:
    
RefreshPercentageThreshold(float64)
    Accepts any float between 0 and 1 (exclusive) that defines the percentage of token validity when to schedule background refreshing

WarningPercentageThreshold(float64)
    Accepts any float between 0 (exclusive) and 1 (inclusive) that defines when the library starts logging warnings that tokens will, eventually expire. This can happen if, for example, the background refresh is failing. It should be higher than the refresh threshold.
    
UserCredentialsProvider(user.CredentialsProvider)
    Accepts any user.CredentialsProvider instance that will provide user credentials for the password grant type
    
ClientCredentialsProvider(client.CredentialsProvider)
    Accepts any client.CredentialsProvider instance that will provide client credentials for the OAuth calls

Contributing
~~~~~~~~~~~~~~~
go-tokens welcomes contributions and questions from the open-source community. First, drop us a line in the Issues Tracker giving us a heads-up on what you'd like to add, change, or discuss. We'll reply and let you know if we think you should go ahead with a PR.

Related/Similar Libraries
~~~~~~~~~~~~~~~

go-tokens is very similar to `tokens`_ and `python-tokens`_.

.. _tokens: https://github.com/zalando-stups/tokens
.. _python-tokens: https://github.com/zalando-stups/python-tokens
.. _user/static.go: https://github.com/zalando/go-tokens/blob/master/user/static.go
.. _client/static.go: https://github.com/zalando/go-tokens/blob/master/client/static.go
