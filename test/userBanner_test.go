//package test
//
//import (
//	"banner-service/internal/cache"
//	"banner-service/internal/config"
//	"banner-service/internal/handler"
//	"banner-service/internal/handler/model/request"
//	"banner-service/internal/repository"
//	"banner-service/internal/service"
//	"encoding/json"
//	"fmt"
//	"github.com/stretchr/testify/suite"
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"strconv"
//	"testing"
//)
//
//type Suite struct {
//	suite.Suite
//
//	hand  *handler.Handler
//	srvc  *service.Service
//	repo  *repository.Repository
//	cache *cache.Cache
//}
//
//func TestSuite(t *testing.T) {
//	suite.Run(t, new(Suite))
//}
//
//func TestMain(m *testing.M) {
//	rc := m.Run()
//	os.Exit(rc)
//}
//
//func (s *Suite) SetupSuite() {
//	cfg := config.NewConfig("../configs/config.yaml")
//
//	//cfg.DB.Host = "localhost"
//	//cfg.Redis.URL = "localhost:6379"
//
//	pg, err := repository.NewPostgres(cfg.DB)
//	if err != nil {
//		s.FailNow("failed postgres connection: ", err)
//	}
//	defer pg.Close()
//
//	red, err := cache.NewRedis(cfg.Redis)
//	if err != nil {
//		s.FailNow("failed redis connection: ", err)
//	}
//	defer red.Close()
//
//	s.repo = repository.NewRepository(pg)
//	s.cache = cache.NewCache(red)
//	s.srvc = service.NewService(s.repo, s.cache)
//	s.hand = handler.NewHandler(s.srvc)
//
//	if err = s.loadDB(); err != nil {
//		s.FailNow("failed postgres load", err)
//	}
//}
//
//func (s *Suite) loadDB() error {
//	bans := make([]request.Banner, 5)
//
//	for i := 0; i < 5; i++ {
//		bans[i].TagIds = []uint64{uint64(i*3 + 1), uint64(i*3 + 2), uint64(i*3 + 3)}
//		bans[i].FeatureId = new(uint64)
//		*bans[i].FeatureId = uint64(i + 1)
//		bans[i].Content = new(json.RawMessage)
//		*bans[i].Content = json.RawMessage([]byte("{\"title\":" + strconv.Itoa(i) + "}"))
//		bans[i].IsActive = new(bool)
//		*bans[i].IsActive = true
//		if _, err := s.repo.PostBanner(bans[i]); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func (s *Suite) TestGetUserBannerFromDB() {
//
//	req := httptest.NewRequest("GET", "/user_banner?tag_id=1&feature_id=1&use_last_revision=false", nil)
//	req.Header.Set("accept", "application/json")
//	req.Header.Set("token", "admin_token")
//
//	w := httptest.NewRecorder()
//
//	s.hand.UserBanner(w, req)
//
//	r := s.Require()
//	body, _ := ioutil.ReadAll(w.Result().Body)
//
//	fmt.Println(string(body))
//
//	r.Equal(http.StatusOK, w.Result().StatusCode)
//}
