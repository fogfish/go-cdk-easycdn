package easycdn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fogfish/go-cdk-easycdn/easycdn/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/fogfish/go-cdk-easycdn/easycdn/internal"
)

type Cdn interface {
	constructs.Construct
	Distribution() awscloudfront.Distribution
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Cdn
type jsiiProxy_Cdn struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Cdn) Distribution() awscloudfront.Distribution {
	var returns awscloudfront.Distribution
	_jsii_.Get(
		j,
		"distribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdn) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCdn(scope constructs.Construct, id *string, props *CdnProps) Cdn {
	_init_.Initialize()

	if err := validateNewCdnParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cdn{}

	_jsii_.Create(
		"easycdn.Cdn",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCdn_Override(c Cdn, scope constructs.Construct, id *string, props *CdnProps) {
	_init_.Initialize()

	_jsii_.Create(
		"easycdn.Cdn",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Cdn_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdn_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"easycdn.Cdn",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cdn) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

