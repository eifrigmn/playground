# golang单元测试框架调研

单元测试是用来对一个函数，一个模块或者一个类来进行正确性检验的工具，开发人员可以通过单元测试来测试项目中的函数、功能模块等并确保其按预期运行，以此降低bug生成率，并提高开发效率。

根据不同的需求，可以采用不同的单元测试框架。

## testing

~~~go
func TestPalindrome(t *testing.T) {
        // "detartrated"是一个回文字符串，因此IsPalindrome("detartrated")的返回值应该为true
        // 如果返回false，则表示其实现有问题，需要使用t.Error进行错误报告
	if !IsPalindrome("detartrated") {
	    t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
	    t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
    if IsPalindrome("palindrome") {
        t.Error(`IsPalindrome("palindrome") = true`) 
    }
}
~~~

使用`testing`可以完成大部分单元测试需求，而为了更优雅及高效地完成单元测试，需要借助一些第三方包。

## 断言

为了判断被测试模块的执行效果，我们常常会引入冗杂的`if...else...`的结构，显得很不优雅，并且难于阅读。因此需要引入`断言`来实现这一功能。这里引入两个第三方单元测试框架：

[testify](https://github.com/stretchr/testify)

~~~go
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
  assert.False(t, IsPalindrome("detartrated"))
  assert.True(t, IsPalindrome("kayak"))
}
~~~

[goconvey](github.com/smartystreets/goconvey)

~~~go
import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestPalindrome(t *testing.T) {
  Convey("test func Palindrome", t, func() {
		So(IsPalindrome("detartrated"), ShouldBeFalse)
    So(IsPalindrome("kayak"), ShouldBeTrue)
	})
}
~~~

## 打桩

实际的单元测试中，某个功能模块往往会有很多的依赖项：

数据库连接、文件I/O、其他函数模块、全局变量等，为了专注于主要对象的测试，我们可以使用一个模拟对象来代替次要模块，以简化测试。

例如一个典型的读配置文件操作：

```go
type Config struct {
   Name string
   Data string
}
func LoadConfig(path string)(*Config, error) {
   f, err := os.Open(path)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   var cfgByte []byte
   _,err = f.Read(cfgByte)
   if err != nil {
      return nil, err
   }
   var cfg Config
   err = json.Unmarshal(cfgByte, &cfg)
   return &cfg, err

}

func (c *Config)GetConfigName() string {
   return c.Name
}
```

### [gomonkey](https://github.com/agiledragon/gomonkey)

gomonkey实现了单元测试中的猴子补丁，可以很方便地为方法、全局变量打桩，同时可以指定行为序列。

以读取上文提到的config文件为例，我们的目的是为了测试`GetConfigName`方法，而与之弱相关的配置文件读取操作(LoadConfig)，我们可以使用gomonkey进行打桩：

~~~go
func TestConfig_GetConfigName(t *testing.T) {
	p := ApplyFunc(LoadConfig, func(_ string)(*Config,error){
		return &Config{
			Name: "mock_config",
			Data: "mock data of config",
		}, nil
	})
	defer p.Reset()
	cfg, err := LoadConfig("")
	assert.Equal(t, nil, err)
	assert.Equal(t, "mock_config", cfg.GetConfigName())
}
~~~

### [GoMock](https://github.com/golang/mock)

gomock是官方提供的模拟框架，可以与testing很好的集成。主要针对的是**对象+接口的数据结构**。也就是框架完成了繁琐的实现接口的工作。

例如存在一个接口`Male`

~~~go
type Male interface {
  Get(id int64) error 
}
~~~

结构体`User`基于`Male`实现了一系列方法

~~~go
type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) error {
	return u.Person.Get(id)
}
~~~

为了测试结构体`User`对应的方法，我们可以使用`gomock`mock出一个`Male`接口，将主要的测试精力回归到`User`方法的测试上来

1. 先使用`mockgen`生成对应文件的mock文件

   ~~~shell
   mockgen -source=./person/male.go -destination=./mock/male_mock.go -package=mock
   ~~~

2. 进行测试

   ~~~go
   func TestUser_GetUserInfo(t *testing.T) {
   	ctl := gomock.NewController(t)
   
   	var id int64 = 1
   	mockMale := mock.NewMockMale(ctl)
   	gomock.InOrder(
   		mockMale.EXPECT().Get(id).Return(errors.New("test error")),
   	)
   	user := NewUser(mockMale)
   	err := user.GetUserInfo(id)
   	assert.Equal(t, errors.New("test error"), err)
   }
   ~~~

## Reference

[Go工程化(八) 单元测试 - Mohuishou (lailin.xyz)](https://lailin.xyz/post/go-training-week4-unit-test.html)

[Unit Testing in Go | PullRequest Blog](https://www.pullrequest.com/blog/unit-testing-in-go/#:~:text=Unit testing is very common in Go (Golang),help to steer you in the right direction.)

[如何高效编写Go单元测试（一） - 掘金 (juejin.cn)](https://juejin.cn/post/6908938380114034701)

[如何高效编写Go单元测试（二） - 掘金 (juejin.cn)](https://juejin.cn/post/6916134886022037511)

[gomock 使用 - Silverming (xiaoming.net.cn)](https://xiaoming.net.cn/2021/06/29/gomock 使用/)

