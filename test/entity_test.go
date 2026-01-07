package test

import(
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"testing"
	"unittest/entity"
)

func TestEntity(t *testing.T){
	t.Run("valid",func(t *testing.T){
		g := NewGomegaWithT(t)
		student := entity.Student{
			StudentID:"B1234567",
			Name:"Ho",
			GPA:3.5,
		}
		ok, err := govalidator.ValidateStruct(student)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
	t.Run("studentid false",func(t *testing.T){
		g:= NewGomegaWithT(t)
		student := entity.Student{
			StudentID:"TT12384000",
			Name:"Ho",
			GPA:3.5,
		}
		ok, err := govalidator.ValidateStruct(student)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("StudentID is invalid"))
	})

	t.Run("invalidtwo",func(t *testing.T){
		g := NewGomegaWithT(t)
		student := entity.Student{
			StudentID:"",
			Name:"",
			GPA:4.0,
		}
		ok , err := govalidator.ValidateStruct(student)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name is invalid;StudentID is required"))
	})
}