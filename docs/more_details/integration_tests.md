# Integration tests

In this section we will be looking at several types of integration tests provided by Gosoline. All of them can be used together in the same test, and Gosoline's `suite` package is responsible for reading their configuration, starting and running an application, then running the actual tests cases. Below is their description:

### Base test case

This is the simplest type of Gosoline test case. All it needs is a suite object having a method that starts with `Test`, which has no inputs nor outputs. An example can be found in `more_details/stream-consumer`:

[embedmd]:# (../../examples/more_details/stream-consumer-test/stream_consumer_test.go /func \(s \*ConsumerTestSuite\) TestComponents/ /\n}/)
```go
func (s *ConsumerTestSuite) TestComponents() {
	s3 := s.Env().Component("s3", "default")
	s.NotNil(s3)

	streamInput := s.Env().Component("streamInput", "consumerInput")
	s.NotNil(streamInput)

	streamOutput := s.Env().Component("streamOutput", "publisher-outputEvent")
	s.NotNil(streamOutput)
}
```

This particular test makes use of _suite.Suite_'s methods to get and check if three components have been wired.

### Application test case

_Application_ test cases are just like _Base_ test cases, the only difference being that they make use of an _suite.AppUnderTest_ object. An example can be found in the same file:

[embedmd]:# (../../examples/more_details/stream-consumer-test/stream_consumer_test.go /func \(s \*ConsumerTestSuite\) TestSuccessTwice/ /\n}/)
```go
func (s *ConsumerTestSuite) TestSuccessTwice(app suite.AppUnderTest) {
	consumer := s.Env().StreamInput("consumerInput")
	s.NotNil(consumer)

	consumer.Publish(mdl.Box(uint(2)), nil)
	consumer.Publish(mdl.Box(uint(3)), nil)

	app.Stop()
	app.WaitDone()

	var result int
	s.Env().StreamOutput("publisher-outputEvent").Unmarshal(0, &result)
	s.Equal(3, result)

	s.Env().StreamOutput("publisher-outputEvent").Unmarshal(1, &result)
	s.Equal(4, result)
}
```

This test publishes two items into the application's stream input, waits until the application is done, then reads its outputs, and compares them with their expected values. 

### ApiServer test case

If you want to make a standard API call and read the response, you can use Gosoline's _ApiServerTestCase_. These type of tests issue HTTP calls and compare their responses against a predefined value, and they need to follow a predefined structure. Below we can see and example:

[embedmd]:# (../../examples/getting_started/integration/api_test.go /func \(s \*ApiTestSuite\) Test_Euro/ /\n}/)
```go
func (s *ApiTestSuite) Test_Euro() *suite.ApiServerTestCase {
	return &suite.ApiServerTestCase{
		Method:             http.MethodGet,
		Url:                "/euro/10/GBP",
		Headers:            map[string]string{},
		ExpectedStatusCode: http.StatusOK,
		Assert: func(response *resty.Response) error {
			result, err := strconv.ParseFloat(string(response.Body()), 64)
			s.NoError(err)
			s.Equal(8.0, result)

			return nil
		},
	}
}
```

Notice the predefined structure of an `ApiServerTestCase`: it is an object that has fields for all the information needed in performing an HTTP call, the expected status code, and a method that tests the result for correctness.

In the example from `examples/getting_started/integration/api_test.go` we see two types of test cases (_ApiServer_ and _ApiServer extended_) in the same test suite. In fact, all types of tests can be part of the same test suite.

### ApiServer extended test case

In the [Integration tests for your API](../getting_started/integration_tests.md) example, we saw the following test:

[embedmd]:# (../../examples/getting_started/integration/api_test.go /func \(s \*ApiTestSuite\) Test_ToEuro/ /\n}/)
```go
func (s *ApiTestSuite) Test_ToEuro(_ suite.AppUnderTest, client *resty.Client) error {
	var result float64

	response, err := client.R().
		SetResult(&result).
		Execute(http.MethodGet, "/euro/10/GBP")

	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode())
	s.Equal(8.0, result)

	return nil
}
```

The `Test_ToEuro` method makes a GET call to an endpoint, in order to receive back the euro exchange value for 10 GBP. Lastly, it checks if the received value is 8.0.

In order to make this GET call `Test_ToEuro` does not need to concern itself about the IP on which the exchange application is running, nor its port. All `Test_ToEuro` needs to know is the URL path of that endpoint, as the _client_ object does the rest.

This _client_ object is provided by Gosoline, whenever at least one of a test suite's methods has the above signature. The main advantage of the _client_ object is that it allows you to control when requests are executed, and gives you access to the endpoint, by providing the host & port. Therefore, for example, if you want to time your requests, this is the way to go.

[embedmd]:# (../../examples/getting_started/integration/api_test.go /func \(s \*ApiTestSuite\) SetupApiDefinitions/ /return definer.ApiDefiner\n}/)
```go
func (s *ApiTestSuite) SetupApiDefinitions() apiserver.Definer {
	return definer.ApiDefiner
}
```

Implementing _SetupApiDefinitions_ when configuring your test will inform Gosoline that an API is being tested, thus the _client_ object will contain meaningful data.

Forgetting to implement _SetupApiDefinitions_, and the trying to run an _Apiserver_, or an _Apiserver extended_ test case, will result in an error reminding you that `the suite has to implement the TestingSuiteApiDefinitionsAware interface`.

### Stream test case

If you have a Gosoline application that takes its input from a stream, you need a test which can run your application locally, send data to a stream, and check your application's output for correctness. 

The stream-consumer application, found in `more_details/stream-consumer`, reads unsigned integers from the `consumerInput` input, increments them by one, and publishes them to the `publisher-outputEvent` output. Below is an extract from its integration test:

[embedmd]:# (../../examples/more_details/stream-consumer-test/stream_consumer_test.go /func \(s \*ConsumerTestSuite\) SetupSuite/ /\n}/)
```go
func (s *ConsumerTestSuite) SetupSuite() []suite.Option {
	return []suite.Option{
		suite.WithLogLevel("debug"),
		suite.WithConfigFile("../stream-consumer/config.dist.yml"),
		suite.WithModule("consumerModule", stream.NewConsumer("uintConsumer", consumer.NewConsumer())),
	}
}
```

It is making use of the same `config.dist.yml` file as `stream-consumer`, and it will be using the module created by _consumer.NewConsumer_. A `StreamTestCase` is very similar to an `ApiServerTestCase`:

[embedmd]:# (../../examples/more_details/stream-consumer-test/stream_consumer_test.go /func \(s \*ConsumerTestSuite\) TestSuccess/ /\n}/)
```go
func (s *ConsumerTestSuite) TestSuccess() *suite.StreamTestCase {
	return &suite.StreamTestCase{
		Input: map[string][]suite.StreamTestCaseInput{
			"consumerInput": {
				{
					Attributes: nil,
					Body:       mdl.Box(uint(5)),
				},
			},
		},
		Assert: func() error {
			var result int
			s.Env().StreamOutput("publisher-outputEvent").Unmarshal(0, &result)

			s.Equal(6, result)

			return nil
		},
	}
}
```

This `StreamTestCase` is an object defining an input and an `Assert` function. Gosoline will run the module, use this `StreamTestCaseInput` as input for it, then run `Assert`. The `Assert` function reads the first element from a stream, _publisher-outputEvent_, then compares it with an expected result. Notice that the input has a key called `consumerInput`, because in this application's `config.dist.yml` file, we have configured an input named `consumerInput`.

Notice how the stream output object was obtained: `s.Env().StreamOutput("publisher-outputEvent")`. In a similar manner, Gosoline's _suite.Suite_ object can provide other useful components: `s.Env().DynamoDb("default").Client()`, `s.Env().MySql("default").`, `s.Env().Redis("default").Client()`, etc.

### Subscriber test case

The subscriber test case is similar to a _StreamTestCase_. It needs a methods whose name starts with `Test`, has a _suite.Suite_ receiver, and returns an `suite.SubscriberTestCase` and an error.

```go
func (s *SubscriberTestSuite) TestSuccess() (suite.SubscriberTestCase, error) {
	return suite.DdbTestCase(suite.DdbSubscriberTestCase{
		Name:          "client",
		SourceModelId: "mcoins.marketing.management.client",
		TargetModelId: "mcoins.marketing.terminal-affiliate-click.client",
		Input: &terminal_affiliate_click.ClientInputV0{
			Id: 42,
			StoreId:   "my.store.id",
		},
		Assert: func(t *testing.T, fetcher *suite.DdbSubscriberFetcher) {
			actual := &terminal_affiliate_click.Client{}
			fetcher.ByHash(uint(42), actual)

			expected := &terminal_affiliate_click.Client{
				Id:        42,
				StoreId:   "my.store.id",
			}

			s.Equal(expected, actual)
		},
	})
}
```

This test will publish an item to an input, stops the application and waits for it to finish, then looks inside a ddd table to see if it was written there.

## Shared environment

One of the options for a Gosoline _suite_ integration test is `suite.WithSharedEnvironment()`. When this option is off, each test case will run in its own environment. For example, the fixtures are being loaded for every test case, and any change to a database or stream will only last during that test case alone. When this option is enabled, the environment is created only once and used by all the test cases, and any change done by one test case will be available to the ones who follow.

## Auto detect components

Another option each _suite_ test offers is `WithoutAutoDetectedComponents`. This simply adds one extra options to the test, which tells it to skip one of the components configured in any potential `config.dist.yml`. The skipped component's name is given as a parameter to `WithoutAutoDetectedComponents`. Also note that while a component can be skipped by auto detect, it can still be added manually to the test via an option.
