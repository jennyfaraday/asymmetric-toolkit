package logger

/*
	TagTracker::Create() will create a new tag using the given tag name
	and return an associated numeric TagId.
*/
import (
	"regexp"
	"strings"
)

func (o *TagTracker) Create(rawTagName string) TagId {

	re := regexp.MustCompile(tagPattern)

	if rawTagName == "" {
		panic("logger/TagTracker::Create() encountered empty tagName")
	}

	o.initTagNames()
	o.initTagIds()
	o.checkSizeLimit()

	id := o.nextTag
	tagName := strings.TrimSpace(rawTagName)

	if re.MatchString(tagName) {
		o.lock.Lock()
		defer o.lock.Unlock()
		o.tagIds[id] = tagName
		o.tagNames[tagName] = id
		o.nextTag++
	} else {
		panic("Invalid Logger/TagName.  Expected:" + tagPattern)
	}
	return id
}

func (o *TagTracker) initTagNames() {
	if o.tagNames == nil {
		o.tagNames = make(TagNameDictionary, 1)
	}
}

func (o *TagTracker) initTagIds() {
	if o.tagIds == nil {
		o.tagIds = make(TagDictionary, 1)
	}
}

func(o *TagTracker) checkSizeLimit(){
	if len(o.tagIds) >= maxTagTrackerDictionarySize {
		panic("Too many tags are currently open.  Create failed.")
	}
}