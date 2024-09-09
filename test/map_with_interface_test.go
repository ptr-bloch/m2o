package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type InterfaceForMap interface {
	GetName() string
}

type interfaceObject1 struct {
	Name string
	F1   *string
}

func (s interfaceObject1) GetName() string {
	return s.Name + "1"
}

type interfaceObject2 struct {
	Name string
}

func (s interfaceObject2) GetName() string {
	return s.Name + "2"
}

func TestMapWithInterfaceValuesDecoding(t *testing.T) {
	source := map[string]interface{}{
		"key1": map[string]interface{}{
			"Name": "John",
		},
		"key2": map[string]interface{}{
			"Name": "John",
		},
	}

	profile := m2o.NewProfile()

	decode, err := m2o.NewDecoder(map[string]InterfaceForMap{
		"key1": interfaceObject1{},
		"key2": &interfaceObject2{},
	}, m2o.WithProfile(profile))

	var result map[string]InterfaceForMap

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decode.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result == nil ||
		result["key1"] == nil ||
		result["key2"] == nil ||
		result["key1"].GetName() != "John1" ||
		result["key2"].GetName() != "John2" {
		t.Errorf("Decoding failed: got %+v", result)
	}

	checkMemoryUsage(t, profile, &result)
}

var loremIpsum = "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut eros et nisl sagittis vestibulum. Nullam nulla eros, ultricies sit amet, nonummy id, imperdiet feugiat, pede. Sed lectus. Donec mollis hendrerit risus. Phasellus nec sem in justo pellentesque facilisis. Etiam imperdiet imperdiet orci. Nunc nec neque. Phasellus leo dolor, tempus non, auctor et, hendrerit quis, nisi. Curabitur ligula sapien, tincidunt non, euismod vitae, posuere imperdiet, leo. Maecenas malesuada. Praesent congue erat at massa. sdsd"
var elMap = map[string]interface{}{
	"String": loremIpsum,
	"Images": []string{"image1.png", "image2.jpg"},
}
var source = map[string]interface{}{
	"key2": elMap,
	"key3": elMap,
}
